package csv

import (
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

func TestDeserializeUnknownTeamAbbreviation(t *testing.T) {
	deserializer := TeamAbbreviationDeserializer{
		make(map[string]models.Team),
	}
	team, err := deserializer.deserialize("foo")
	if nil != team || err == nil {
		t.Errorf("expected an error")
	}
}
