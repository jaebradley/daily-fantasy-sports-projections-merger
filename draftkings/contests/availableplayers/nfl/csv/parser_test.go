package csv

import (
	"bytes"
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("Position,Name + ID,Name,ID,Roster Position,Salary,Game Info,TeamAbbrev,AvgPointsPerGame\n")
	buffer.WriteString("RB,Derrick Henry (19551612),Derrick Henry,19551612,RB/FLEX,9000,TEN@JAX 10/10/2021 01:00PM ET,TEN,28.12\n")

	parser := Parser{
		timeDeserializer:     &ContestStartTimeDeserializer{},
		salaryDeserializer:   &SalaryDeserializer{},
		playerIDDeserializer: &PlayerIDDeserializer{},
		contestPositionsDeserializer: &ContestPositionsDeserializer{
			contestPositionDeserializer: &ContestPositionDeserializer{
				positionsByAbbreviation: map[string]models.ContestPosition{
					"RB":   models.RUNNINGBACK,
					"FLEX": models.FLEX,
				},
			},
			separator: '/',
		},
		teamDeserializer: &TeamAbbreviationDeserializer{
			map[string]models.Team{
				"TEN": models.TENNESSEETITANS,
				"JAX": models.JACKSONVILLEJAGUARS,
			},
		},
		opponentDeserializer: &TeamAbbreviationDeserializer{
			map[string]models.Team{
				"TEN": models.TENNESSEETITANS,
				"JAX": models.JACKSONVILLEJAGUARS,
			},
		},
	}

	playerDetails, err := parser.Deserialize(&buffer)
	if nil != err {
		t.Errorf("did not expect an error")
	} else if 1 != len(playerDetails) {
		t.Errorf("invalid output")
	}
}
