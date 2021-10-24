package csv

import (
	"bytes"
	"reflect"
	"testing"

	coreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/models"
	rotogrindersSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/serialization"
)

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("player_id,team,opp,pos,name,fpts,proj_own,smash,minutes,ceil,floor,min_exposure,max_exposure,rg_value,salary,custom,rg_id,partner_id\n")
	buffer.WriteString("457688,SAS,MIL,PF,Al-Farouq Aminu,0,,0,,,,0,,-19.50,3000,0,1279,457688")

	parser := Parser{
		PlayerDeserializer: &rotogrindersSerialization.DefaultPlayerDeserializer{
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
		ID:   457688,
		Name: "Al-Farouq Aminu",
	}
	expectedPlayerDetails[expectedPlayer] = 0

	if false == reflect.DeepEqual(playerDetails, expectedPlayerDetails) {
		t.Errorf("invalid output")
	}
}
