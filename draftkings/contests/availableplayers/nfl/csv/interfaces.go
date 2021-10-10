package csv

import (
  "os"
)

// Deserializer deserializes a File into a mapping of player to the player's contest details
type Deserializer interface {
  deserialize(*os.File) (map[Player]NFLContestPlayerDetails, error)
}

