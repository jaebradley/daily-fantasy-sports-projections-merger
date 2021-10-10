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
  reader *Reader
  timeDeserializer *TimeDeserializer
  salaryDeserializer *SalaryDeserializer
  playerIDDeserializer *PlayerIDDeserializer
  playerNameDeserializer *PlayerNameDeserializer
  contestPositionDeserializer *ContestPositionDeserializer
  teamDeserializer *TeamDeserializer
}

func (p Parser) deserialize(*File) (map[Player]NFLContestPlayerDetails, error) {
  detailsByPlayer := make(map[Player]NFLContestPlayerDetails)

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
  }

  return detailsByPlayer
}

