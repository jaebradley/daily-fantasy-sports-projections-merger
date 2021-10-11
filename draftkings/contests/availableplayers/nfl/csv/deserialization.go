package csv

import (
	"fmt"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

// TeamAbbreviationDeserializer deserializes abbreviations to team
type TeamAbbreviationDeserializer struct {
	teamsByAbbreviation map[string]models.Team
}

// Deserialize team abbreviations
func (d *TeamAbbreviationDeserializer) Deserialize(abbreviation string) (*models.Team, error) {
	team, exist := d.teamsByAbbreviation[abbreviation]
	if exist {
		return &team, nil
	}

	return nil, fmt.Errorf("No team identified for abbreviation %s", abbreviation)
}

// ContestPositionDeserializer deserializes abbreviations to contest position
type ContestPositionDeserializer struct {
	positionsByAbbreviation map[string]models.ContestPosition
}

// Deserialize contest position abbreviations
func (d *ContestPositionDeserializer) Deserialize(abbreviation string) (*models.ContestPosition, error) {
	position, exist := d.positionsByAbbreviation[abbreviation]
	if exist {
		return &position, nil
	}

	return nil, fmt.Errorf("No position identified for value: %s", abbreviation)
}
