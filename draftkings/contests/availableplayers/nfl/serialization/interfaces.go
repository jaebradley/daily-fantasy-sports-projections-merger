package serialization

import (
  "time"
  "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

// TeamDeserializer deserializes a string into a team
type TeamDeserializer interface {
  deserialize(string) (Team, error)
}

// ContestPositionDeserializer deserializes a string into a ContestPosition
type ContestPositionDeserializer interface {
  deserialize(string) (ContestPosition, error)
}

// TimeDeserializer deserializes a string into a time.Time
type TimeDeserializer interface {
  deserialize(string) (time.Time, error)
}

// SalaryDeserializer deserializes a string into a salary
type SalaryDeserializer interface {
  deserialize(string) (float, error)
}

// PlayerIDDeserializer deserializes a string into a Player Id
type PlayerIDDeserializer interface {
  deserialize(string) (int, error)
}

// PlayerNameDeserializer deserializes a string into a Player Name
type PlayerNameDeserializer interface {
  deserialize(string) (string, error)
}
