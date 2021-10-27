package csv

import (
	"fmt"
	"strings"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/serialization"
)

// TeamAbbreviationDeserializer deserializes abbreviations to team
type TeamAbbreviationDeserializer struct {
	TeamsByAbbreviation map[string]models.Team
}

// Deserialize team abbreviations
func (d *TeamAbbreviationDeserializer) Deserialize(abbreviation string) (*models.Team, error) {
	team, exist := d.TeamsByAbbreviation[abbreviation]
	if exist {
		return &team, nil
	}

	return nil, fmt.Errorf("No team identified for abbreviation %s", abbreviation)
}

// ContestPositionDeserializer deserializes abbreviations to contest position
type ContestPositionDeserializer struct {
	PositionsByAbbreviation map[string]models.ContestPosition
}

// Deserialize contest position abbreviations
func (d *ContestPositionDeserializer) Deserialize(abbreviation string) (*models.ContestPosition, error) {
	position, exist := d.PositionsByAbbreviation[abbreviation]
	if exist {
		return &position, nil
	}

	return nil, fmt.Errorf("No position identified for value: %s", abbreviation)
}

// ContestPositionsDeserializer deserailizes positions using a separator and a position deserializer
type ContestPositionsDeserializer struct {
	PositionDeserializer serialization.ContestPositionDeserializer
	Separator            rune
}

// Deserialize deserializes a value containing many positions into a mapping of position to order in which thy were defined
func (d *ContestPositionsDeserializer) Deserialize(value string) (map[models.ContestPosition]int, error) {
	indicesByContestPosition := make(map[models.ContestPosition]int)
	parts := strings.Split(value, string(d.Separator))
	for index, part := range parts {
		position, err := d.PositionDeserializer.Deserialize(part)
		if nil != err {
			return nil, err
		}
		indicesByContestPosition[*position] = index
	}
	return indicesByContestPosition, nil
}
