package csv

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/serialization"
)

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("DFS ID,Name,Pos,Team,Opp,Status,Salary,Actual,SS Proj,My Proj,Value,SS Own,My Own,Min Exp,Max Exp,dk_points,dk_25_percentile,dk_50_percentile,dk_75_percentile,dk_85_percentile,dk_95_percentile,dk_99_percentile,fd_points,fd_25_percentile,fd_50_percentile,fd_75_percentile,fd_85_percentile,fd_95_percentile,fd_99_percentile,yahoo_points,yahoo_25_percentile,yahoo_50_percentile,yahoo_75_percentile,yahoo_85_percentile,yahoo_95_percentile,yahoo_99_percentile,dk_std,fd_std,yahoo_std,PTS,Min,2PT,3PT,RB,Off Reb,Def Reb,AST,STL,BLK,TO\n")
	buffer.WriteString("19760145,Nikola Jokic,\"C,UTIL\",DEN,PHX,,10400,,49.9541,49.95,4.802884615384616,17.3,17.3,,,49.9541,42.25,49.5,57.25,61.75,69,77.75,47.8057,40.7,47.6,54.9,59,66.305,74.401,47.8057,40.7,47.6,54.9,59,66.305,74.401,11.4371,11.4371,10.9266,25.202,32.855,9.58067,1.164,10.057,2.29167,7.76533,6.23067,0.664,0.527333,2.38467\n")

	parser := Parser{
		PlayerDeserializer: &serialization.DefaultPlayerDeserializer{
			IdDeserializer:   &serialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	playerDetails, err := parser.Deserialize(&buffer)
	if nil != err {
		t.Errorf("did not expect an error: %q", err)
	} else if 1 != len(playerDetails) {
		t.Errorf("invalid output")
	}

	expectedPlayerDetails := make(map[models.Player]float64)
	expectedPlayer := models.Player{
		Id:   19760145,
		Name: "Nikola Jokic",
	}
	expectedPlayerDetails[expectedPlayer] = 49.9541

	if false == reflect.DeepEqual(playerDetails, expectedPlayerDetails) {
		t.Errorf("invalid output")
	}
}
