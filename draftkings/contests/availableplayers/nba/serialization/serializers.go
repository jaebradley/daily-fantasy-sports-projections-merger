package serialization

import (
	"strings"
	"sync"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
)

var once sync.Once

var abbreviationsByContestPosition = map[models.ContestPosition]string{
	models.POINTGUARD:    "PG",
	models.SHOOTINGGUARD: "SG",
	models.SMALLFORWARD:  "SF",
	models.POWERFORWARD:  "PF",
	models.CENTER:        "C",
	models.GUARD:         "G",
	models.FORWARD:       "F",
	models.UTILITY:       "UTIL",
}

var defaultContestPositionsSerializer *DefaultContestPositionsSerializer

type ContestPositionsSerializer interface {
	Serialize(positions []models.ContestPosition) string
}

type ContestPositionSerializer interface {
	Serialize(position models.ContestPosition) string
}

type DefaultContestPositionSerializer struct {
	AbbreviationsByContestPosition map[models.ContestPosition]string
}

func (s *DefaultContestPositionSerializer) Serialize(position models.ContestPosition) string {
	abbreviation, exist := s.AbbreviationsByContestPosition[position]
	if exist {
		return abbreviation
	}
	return ""
}

type DefaultContestPositionsSerializer struct {
	PositionSerializer ContestPositionSerializer
	Separator          rune
}

func (s *DefaultContestPositionsSerializer) Serialize(positions []models.ContestPosition) string {
	serialized := make([]string, len(positions))
	for index, position := range positions {
		serialized[index] = s.PositionSerializer.Serialize(position)
	}

	return strings.Join(serialized, string(s.Separator))
}

func GetDefaultContestPositionsSerializer() *DefaultContestPositionsSerializer {
	once.Do(func() {
		defaultContestPositionsSerializer = &DefaultContestPositionsSerializer{
			PositionSerializer: &DefaultContestPositionSerializer{
				AbbreviationsByContestPosition: abbreviationsByContestPosition,
			},
			Separator: '/',
		}
	})
	return defaultContestPositionsSerializer
}
