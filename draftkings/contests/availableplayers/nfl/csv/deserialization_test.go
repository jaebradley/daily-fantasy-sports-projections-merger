package csv

import (
	"fmt"
	"reflect"
	"testing"
	"time"

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

func TestDeserializeContestPositions(t *testing.T) {
	contestPositionDeserializer := ContestPositionDeserializer{
		positionsByAbbreviation: map[string]models.ContestPosition{
			"FOO": models.QUARTERBACK,
			"BAR": models.RUNNINGBACK,
		},
	}

	deserializer := ContestPositionsDeserializer{
		contestPositionDeserializer: &contestPositionDeserializer,
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

func TestDeserializeValidContestStartTime(t *testing.T) {
	easternTimeZone, err := time.LoadLocation("America/New_York")
	if nil != err {
		t.Errorf("could not load eastern time zone")
	}
	deserializer := ContestStartTimeDeserializer{}
	startTime, err := deserializer.Deserialize("TEN@JAX 10/10/2021 01:00PM ET")
	if nil != err {
		fmt.Println(err)
		t.Errorf("expected no error")
	} else if !startTime.Equal(time.Date(2021, 10, 10, 13, 0, 0, 0, easternTimeZone)) {
		t.Errorf("unexpected time value")
	}
}
