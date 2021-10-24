package csv

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
)

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("\"Name\",\"Fpts\",\"Position\",\"Team\",\"Opponent\",\"Minutes\",\"Salary\",\"Pts/$\",\"Value\"\n")
	buffer.WriteString("\"Stephen Curry\",\"56.41\",\"PG\",\"GSW\",\"SAC\",\"37.9\",\"11,000\",\"5.13\",\"1.41\"\n")

	parser := Parser{
		PlayerNameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	playerDetails, err := parser.Deserialize(&buffer)
	if nil != err {
		t.Errorf("did not expect an error: %q", err)
	} else if 1 != len(playerDetails) {
		t.Errorf("invalid output")
	}

	expectedPlayerDetails := make(map[string]float64)
	expectedPlayerDetails["Stephen Curry"] = 56.41

	if false == reflect.DeepEqual(playerDetails, expectedPlayerDetails) {
		t.Errorf("invalid output")
	}
}
