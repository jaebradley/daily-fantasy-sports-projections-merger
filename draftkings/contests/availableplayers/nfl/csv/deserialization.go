package csv

import (
  "errors"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/serialization"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)


// TeamAbbreviationDeserializer deserializes abbreviations to team
type TeamAbbreviationDeserializer struct {
  teamsByAbbreviation map[string]models.Team
}

func (d *TeamAbbreviationDeserializer) deserialize(string abbreviation) (*models.Team, error) {
  team, exist := d.teamsByAbbreviation[abbreviation]
  if exist {
    return &team, nil
  }

  return nil, errors.New(fmt.Printf("No team identified for abbreviation %b", abbreviation))
}
