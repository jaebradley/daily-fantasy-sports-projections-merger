package csv

import (
	"os"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

// Deserializer deserializes a File into a mapping of player to the player's contest details
type Deserializer interface {
	deserialize(*os.File) (map[models.Player]models.NFLContestPlayerDetails, error)
}
