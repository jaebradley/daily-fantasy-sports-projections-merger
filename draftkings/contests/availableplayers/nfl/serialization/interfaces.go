package serialization

import (
  "time"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

// TeamDeserializer deserializes a string into a team
type TeamDeserializer interface {
  deserialize(value string) (*models.Team, error)
}

// ContestPositionDeserializer deserializes a string into a ContestPosition
type ContestPositionDeserializer interface {
  deserialize(value string) (*models.ContestPosition, error)
}

// ContestPositionsDeserializer deserializes a string into multiple ContestPositions
type ContestPositionsDeserializer interface {
  deserialize(value string) (map[models.ContestPosition]int, error)
}

// TimeDeserializer deserializes a string into a time.Time
type TimeDeserializer interface {
  deserialize(value string) (*time.Time, error)
}

// SalaryDeserializer deserializes a string into a salary
type SalaryDeserializer interface {
  deserialize(value string) (float, error)
}

// PlayerIDDeserializer deserializes a string into a Player Id
type PlayerIDDeserializer interface {
  deserialize(value string) (int, error)
}

// PlayerNameDeserializer deserializes a string into a Player Name
type PlayerNameDeserializer interface {
  deserialize(value string) (string, error)
}
