package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/serialization"
)

// Parser parses CSV files into a mapping of player to nfl contest player details
type Parser struct {
	reader                       csv.Reader
	timeDeserializer             serialization.TimeDeserializer
	salaryDeserializer           serialization.SalaryDeserializer
	playerIDDeserializer         serialization.PlayerIDDeserializer
	playerNameDeserializer       serialization.PlayerNameDeserializer
	contestPositionsDeserializer serialization.ContestPositionsDeserializer
	teamDeserializer             serialization.TeamDeserializer
	opponentDeserializer         serialization.TeamDeserializer
}

func (p *Parser) deserialize(*os.File) (map[models.Player]models.NFLContestPlayerDetails, error) {
	detailsByPlayer := make(map[models.Player]models.NFLContestPlayerDetails)

	_, err := p.reader.Read()

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

		playerID, err := p.playerIDDeserializer.Deserialize(record[3])
		if err != nil {
			return nil, err
		}

		playerName, err := p.playerNameDeserializer.Deserialize(record[2])
		if err != nil {
			return nil, err
		}

		player := models.Player{
			Id:   playerID,
			Name: playerName,
		}

		_, exist := detailsByPlayer[player]
		if exist {
			return nil, errors.New("duplicate player found")
		}

		team, err := p.teamDeserializer.Deserialize(record[8])
		if err != nil {
			return nil, err
		}

		salary, err := p.salaryDeserializer.Deserialize(record[5])
		if err != nil {
			return nil, err
		}

		indexByPosition, err := p.contestPositionsDeserializer.Deserialize(record[4])
		if err != nil {
			return nil, err
		}

		startTime, err := p.timeDeserializer.Deserialize(record[6])
		if err != nil {
			return nil, err
		}

		opponent, err := p.opponentDeserializer.Deserialize(record[6])
		if err != nil {
			return nil, err
		}

		elibigilityByPositions := make(map[models.ContestPosition]bool)
		for contestPosition := range models.ContestPositions {
			_, exist := indexByPosition[contestPosition]
			elibigilityByPositions[contestPosition] = exist
		}

		details := models.NFLContestPlayerDetails{
			Team:                   team,
			Opponent:               opponent,
			StartTime:              startTime,
			Salary:                 salary,
			EligibilityByPositions: elibigilityByPositions,
		}

		detailsByPlayer[player] = details
	}

	return detailsByPlayer, nil
}
