package csv

import (
	"bytes"
	"reflect"
	"testing"

	coreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/models"
	dailyrotoSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/serialization"
)

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("Team,Opp,Player,Slate PlayerID,,Minutes,Usage Rate,Rebound Rate,AssistRate,FanDuel:,Pos,Public %,Optimal%,Leverage,Salary,Points,Value,DraftKings:,Pos,Public %,Optimal%,Leverage,Salary,Points,Value\n")
	buffer.WriteString("PHI,NO,Joel Embiid,0,,33.1,34.49,18.85,21.82,,C,18.8,26.3,7.5,11000,54.04,6.64,,C,20.2,20.6,0.40000000000000213,10600,55.33,5.93\n")

	parser := Parser{
		PlayerDeserializer: &dailyrotoSerialization.DefaultPlayerDeserializer{
			IDDeserializer:   &coreSerialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &coreSerialization.DefaultPlayerNameDeserializer{},
		},
		ProjectionDeserializer: &coreSerialization.DefaultProjectionDeserializer{},
	}

	playerDetails, err := parser.Deserialize(&buffer)
	if nil != err {
		t.Errorf("did not expect an error: %q", err)
	} else if 1 != len(playerDetails) {
		t.Errorf("invalid output")
	}

	expectedPlayerDetails := make(map[models.Player]float64)
	expectedPlayer := models.Player{
		ID:   0,
		Name: "Joel Embiid",
	}
	expectedPlayerDetails[expectedPlayer] = 55.33

	if false == reflect.DeepEqual(playerDetails, expectedPlayerDetails) {
		t.Errorf("invalid output")
	}
}
