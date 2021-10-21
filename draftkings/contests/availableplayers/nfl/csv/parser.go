package csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/serialization"
)

// Parser parses CSV files into a mapping of player to nfl contest player details
type Parser struct {
	timeDeserializer             serialization.TimeDeserializer
	salaryDeserializer           serialization.SalaryDeserializer
	playerIDDeserializer         serialization.PlayerIDDeserializer
	contestPositionsDeserializer serialization.ContestPositionsDeserializer
	teamDeserializer             serialization.TeamDeserializer
	opponentDeserializer         serialization.TeamDeserializer
}

func (p *Parser) Deserialize(bytes *bytes.Buffer) (map[models.Player]models.NFLContestPlayerDetails, error) {
	detailsByPlayer := make(map[models.Player]models.NFLContestPlayerDetails)
	reader := csv.NewReader(bytes)

	_, err := reader.Read()

	if err == io.EOF {
		return detailsByPlayer, nil
	}

	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()

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

		player := models.Player{
			Id:   playerID,
			Name: record[2],
		}

		_, exist := detailsByPlayer[player]
		if exist {
			return nil, errors.New("duplicate player found")
		}

		team, err := p.teamDeserializer.Deserialize(record[7])
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

		elibigilityByPositions := make(map[models.ContestPosition]bool)
		for contestPosition := range models.ContestPositions {
			_, exist := indexByPosition[contestPosition]
			elibigilityByPositions[contestPosition] = exist
		}

		details := models.NFLContestPlayerDetails{
			Team: team,
			// TODO:@jaebradley actually parse opponent
			Opponent:               nil,
			StartTime:              startTime,
			Salary:                 salary,
			EligibilityByPositions: elibigilityByPositions,
		}

		detailsByPlayer[player] = details
	}

	return detailsByPlayer, nil
}
