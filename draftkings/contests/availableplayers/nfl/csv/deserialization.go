package csv

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/serialization"
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

// ContestPositionsDeserializer deserailizes positions using a separator and a position deserializer
type ContestPositionsDeserializer struct {
	contestPositionDeserializer serialization.ContestPositionDeserializer
	separator                   rune
}

// Deserialize deserializes a value containing many positions into a mapping of position to order in which thy were defined
func (d *ContestPositionsDeserializer) Deserialize(value string) (map[models.ContestPosition]int, error) {
	indicesByContestPosition := make(map[models.ContestPosition]int)
	parts := strings.Split(value, string(d.separator))
	for index, part := range parts {
		position, err := d.contestPositionDeserializer.Deserialize(part)
		if nil != err {
			return nil, err
		}
		indicesByContestPosition[*position] = index
	}
	return indicesByContestPosition, nil
}

type ContestStartTimeDeserializer struct {
}

func (d *ContestStartTimeDeserializer) Deserialize(value string) (*time.Time, error) {
	parts := strings.Split(value, " ")
	if 4 != len(parts) {
		return nil, errors.New("expected exactly 4 parts")
	}

	if "ET" != parts[3] {
		return nil, errors.New("unexpected time zone")
	}

	easternTimeZone, err := time.LoadLocation("America/New_York")
	if nil != err {
		return nil, errors.New("unable to load Eastern Time Zone")
	}

	parsedTime, err := time.ParseInLocation("01/02/2006 03:04PM", strings.Join(parts[1:3], " "), easternTimeZone)
	if nil != err {
		return nil, err
	}
	return &parsedTime, nil
}
