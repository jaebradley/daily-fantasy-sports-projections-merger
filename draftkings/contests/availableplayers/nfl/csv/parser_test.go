package csv

import (
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

func TestParser(t *testing.T) {
	parser := Parser{
		reader:               Reader,
		timeDeserializer:     TimeDeserializer{},
		salaryDeserializer:   SalaryDeserializer{},
		playerIDDeserializer: PlayerIDDeserializer{},
		contestPositionsDeserializer: ContestPositionsDeserializer{
			contestPositionDeserializer: &ContestPositionDeserializer{
				positionsByAbbreviation: map[string]models.ContestPosition{
					"FOO": models.QUARTERBACK,
					"BAR": models.RUNNINGBACK,
				},
			},
			separator: ',',
		},
		teamDeserializer: TeamAbbreviationDeserializer{
			map[string]models.Team{
				"foo": models.NEWENGLANDPATRIOTS,
			},
		},
		opponentDeserializer: TeamAbbreviationDeserializer{
			map[string]models.Team{
				"foo": models.NEWENGLANDPATRIOTS,
			},
		},
	}
}
