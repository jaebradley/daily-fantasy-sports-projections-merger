package serialization

import (
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
)

// TeamDeserializer deserializes a string into a team
type TeamDeserializer interface {
	Deserialize(value string) (*models.Team, error)
}

// ContestPositionDeserializer deserializes a string into a ContestPosition
type ContestPositionDeserializer interface {
	Deserialize(value string) (*models.ContestPosition, error)
}

// ContestPositionsDeserializer deserializes a string into multiple ContestPositions
type ContestPositionsDeserializer interface {
	Deserialize(value string) (map[models.ContestPosition]int, error)
}
