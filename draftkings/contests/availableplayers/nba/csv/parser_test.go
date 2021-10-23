package csv

import (
	"bytes"
	"testing"

	coreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/serialization"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
	nbaModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
)

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("Position,Name + ID,Name,ID,Roster Position,Salary,Game Info,TeamAbbrev,AvgPointsPerGame\n")
	buffer.WriteString("C,Joel Embiid (19760143),Joel Embiid,19760143,C/UTIL,10600,PHI@NOP 10/20/2021 08:00PM ET,PHI,50.55")

	parser := Parser{
		TimeDeserializer:   &coreSerialization.DefaultContestStartTimeDeserializer{},
		SalaryDeserializer: &coreSerialization.DefaultSalaryDeserializer{},
		PlayerDeserializer: &coreSerialization.DefaultPlayerDeserializer{
			IdDeserializer:   &coreSerialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &coreSerialization.DefaultPlayerNameDeserializer{},
		},
		ContestPositionsDeserializer: &ContestPositionsDeserializer{
			PositionDeserializer: &ContestPositionDeserializer{
				PositionsByAbbreviation: map[string]nbaModels.ContestPosition{
					"C":    models.CENTER,
					"UTIL": models.UTILITY,
				},
			},
			Separator: '/',
		},
		TeamDeserializer: &TeamAbbreviationDeserializer{
			map[string]nbaModels.Team{
				"PHI": nbaModels.PHILADELPHA76ERS,
			},
		},
		OpponentDeserializer: &TeamAbbreviationDeserializer{
			map[string]nbaModels.Team{
				"PHI": nbaModels.PHILADELPHA76ERS,
			},
		},
	}

	playerDetails, err := parser.Deserialize(&buffer)
	if nil != err {
		t.Errorf("did not expect an error: %q", err)
	} else if 1 != len(playerDetails) {
		t.Errorf("invalid output")
	}
}
