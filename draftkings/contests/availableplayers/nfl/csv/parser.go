package csv

import (
  "errors"
  "os"
  "io"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/csv"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/serialization"
)


// Parser parses CSV files into a mapping of player to nfl contest player details
type Parser struct {
  reader *csv.Reader
  timeDeserializer *serialization.TimeDeserializer
  salaryDeserializer *serialization.SalaryDeserializer
  playerIDDeserializer *serialization.PlayerIDDeserializer
  playerNameDeserializer *serialization.PlayerNameDeserializer
  contestPositionsDeserializer *serialization.ContestPositionsDeserializer
  teamDeserializer *serialization.TeamDeserializer
  opponentDeserializer *serialization.OpponentDeserializer
}

func (p *Parser) deserialize(*os.File) (map[models.Player]models.NFLContestPlayerDetails, error) {
  detailsByPlayer := make(map[models.Player]models.NFLContestPlayerDetails)

  header, err := p.reader.Read()

  if err == io.EOF {
    return detailsByPlayer, nil
  }

  if err != nil {
    return nil, err
  }

  for {
    record, err := p.reader.Read()

    if err == io.EOF {
    return detailsByPlayer, nil
    }

    if err != nil {
      return nil, err
    }

    playerID, err := p.playerIDDeserializer.deserialize(record[3])
    if err != nil {
      return nil, err
    }

    playerName, err := p.playerNameDeserializer.deserialize(record[2])
    if err != nil {
      return nil, err
    }

    player := Player{
      playerID,
      playerName,
    }

    value, exist := detailsByPlayer[player]
    if exist {
      return nil, errors.New("duplicate player found")
    }

    team, err := p.teamDeserializer.deserialize(record[8])
    if err != nil {
      return nil, err
    }

    salary, err := p.salaryDeserializer.deserialize(record[5])
    if err != nil {
      return nil, err
    }

    indexByPosition, err := p.contestPositionsDeserializer.deserialize(record[4])
    if err != nil {
      return nil, err
    }

    startTime, err := p.timeDeserializer.deserialize(record[6])
    if err != nil {
      return nil, err
    }

    opponent, err := p.opponentDeserializer.deserialize(record[6])
    if err != nil {
      return nil, err
    }

    elibigilityByPositions := make(map[models.ContestPosition]bool)
    for contestPosition := range models.ContestPositions {
      value, exist := indexByPosition[contestPosition]
      eligibilityByPositions[contestPosition] = exist
    }

    details := models.NFLContestPlayerDetails{
      team,
      opponent,
      startTime,
      salary,
      eligibilityByPositions,
    }

    detailsByPlayer[player] = details
  }

  return detailsByPlayer
}

