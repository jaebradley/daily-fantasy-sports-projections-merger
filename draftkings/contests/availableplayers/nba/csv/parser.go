package csv

import (
	"encoding/csv"
	"errors"
	"io"

	coreModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/models"
	coreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/serialization"
	nbaModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
	nbaSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/serialization"
)

// Parser parses CSV files into a mapping of player to nba contest player details
type Parser struct {
	TimeDeserializer             coreSerialization.TimeDeserializer
	SalaryDeserializer           coreSerialization.SalaryDeserializer
	PlayerDeserializer           coreSerialization.PlayerDeserializer
	ContestPositionsDeserializer nbaSerialization.ContestPositionsDeserializer
	TeamDeserializer             nbaSerialization.TeamDeserializer
	OpponentDeserializer         nbaSerialization.TeamDeserializer
}

// Deserialize deserializes bytes to a mapping of contest details to a player
func (p *Parser) Deserialize(r io.Reader) (map[coreModels.Player]nbaModels.NBAContestPlayerDetails, error) {
	detailsByPlayer := make(map[coreModels.Player]nbaModels.NBAContestPlayerDetails)
	reader := csv.NewReader(r)

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

		player, err := p.PlayerDeserializer.Deserialize(record[3], record[2])
		if err != nil {
			return nil, err
		}

		_, exist := detailsByPlayer[*player]
		if exist {
			return nil, errors.New("duplicate player found")
		}

		team, err := p.TeamDeserializer.Deserialize(record[7])
		if err != nil {
			return nil, err
		}

		salary, err := p.SalaryDeserializer.Deserialize(record[5])
		if err != nil {
			return nil, err
		}

		indexByPosition, err := p.ContestPositionsDeserializer.Deserialize(record[4])
		if err != nil {
			return nil, err
		}

		startTime, err := p.TimeDeserializer.Deserialize(record[6])
		if err != nil {
			return nil, err
		}

		elibigilityByPositions := make(map[nbaModels.ContestPosition]bool)
		for contestPosition := range nbaModels.ContestPositions {
			_, exist := indexByPosition[contestPosition]
			elibigilityByPositions[contestPosition] = exist
		}

		details := nbaModels.NBAContestPlayerDetails{
			Team: team,
			// TODO:@jaebradley actually parse opponent
			Opponent:               nil,
			StartTime:              startTime,
			Salary:                 salary,
			EligibilityByPositions: elibigilityByPositions,
		}

		detailsByPlayer[*player] = details
	}

	return detailsByPlayer, nil
}
