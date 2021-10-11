package csv

import (
	"fmt"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

// TeamAbbreviationDeserializer deserializes abbreviations to team
type TeamAbbreviationDeserializer struct {
	teamsByAbbreviation map[string]models.Team
}

func (d *TeamAbbreviationDeserializer) deserialize(abbreviation string) (*models.Team, error) {
	team, exist := d.teamsByAbbreviation[abbreviation]
	if exist {
		return &team, nil
	}

	return nil, fmt.Errorf("No team identified for abbreviation %s", abbreviation)
}
