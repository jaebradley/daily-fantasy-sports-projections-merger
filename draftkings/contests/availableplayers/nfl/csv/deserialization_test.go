package csv

import (
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
)

func TestDeserializeUnknownTeamAbbreviation(t *testing.T) {
	deserializer := TeamAbbreviationDeserializer{
		make(map[string]models.Team),
	}
	team, err := deserializer.Deserialize("foo")
	if nil != team || err == nil {
		t.Errorf("expected an error")
	}
}

func TestDeserializeKnownTeamAbbreviation(t *testing.T) {
	deserializer := TeamAbbreviationDeserializer{
		map[string]models.Team{
			"foo": models.NEWENGLANDPATRIOTS,
		},
	}

	team, err := deserializer.Deserialize("foo")
	if *team != models.NEWENGLANDPATRIOTS || err != nil {
		t.Errorf("unexpected team or error")
	}
}

func TestDeserializeKnownContestPosition(t *testing.T) {
	deserializer := ContestPositionDeserializer{
		make(map[string]models.ContestPosition),
	}

	position, err := deserializer.Deserialize("foo")
	if nil != position || nil == err {
		t.Errorf("expected position to be nil and error to exist")
	}
}
