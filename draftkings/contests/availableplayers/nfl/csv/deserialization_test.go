package csv

import (
	"reflect"
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nfl/serialization"
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

func TestDeserializeContestPositions(t *testing.T) {
	var contestPositionDeserializer serialization.ContestPositionDeserializer = ContestPositionDeserializer{
		positionsByAbbreviation: map[string]models.ContestPosition{
			"FOO": models.QUARTERBACK,
			"BAR": models.RUNNINGBACK,
		},
	}

	deserializer := ContestPositionsDeserializer{
		contestPositionDeserializer: contestPositionDeserializer,
		separator:                   ',',
	}

	expectedIndicesByContestPosition := map[models.ContestPosition]int{
		models.QUARTERBACK: 0,
		models.RUNNINGBACK: 1,
	}

	indicesByContestPosition, err := deserializer.Deserialize("FOO,BAR")
	if nil != err {
		t.Errorf("expected no errors")
	} else if !reflect.DeepEqual(expectedIndicesByContestPosition, indicesByContestPosition) {
		t.Errorf("expected positions to match")
	}
}
