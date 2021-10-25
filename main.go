package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	awesomeoCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/awesomeo/csv"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	dailyRotoCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/csv"
	dailyrotoSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/serialization"
	rotogrindersCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/csv"
	rotogrindersSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/serialization"
	saberSimCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/csv"
	saberSimSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/serialization"
)

func main() {
	awesomeoFilePath := flag.String("awesomeo", "", "path to Awesomeo projections downloaded as a CSV file")
	dailyRotoFilePath := flag.String("dailyroto", "", "path to DailyRoto projections downloaded as a CSV file")
	rotogrindersFilePath := flag.String("rotogrinders", "", "path to RotoGrinders projections downloaded as a CSV file")
	sabersimFilePath := flag.String("sabersim", "", "path to Sabersim projections downloaded as a CSV file")
	flag.Parse()

	if "" == *awesomeoFilePath {
		os.Exit(255)
	}

	if "" == *dailyRotoFilePath {
		os.Exit(255)
	}

	if "" == *rotogrindersFilePath {
		os.Exit(255)
	}

	awesomeoFile, err := os.Open(*awesomeoFilePath)
	defer awesomeoFile.Close()
	if nil != err {
		os.Exit(255)
	}

	awesomeoParser := awesomeoCsv.Parser{
		PlayerNameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	awesomeoProjections, err := awesomeoParser.Deserialize(bufio.NewReader(awesomeoFile))
	if nil != err {
		os.Exit(255)
	}

	dailyRotoFile, err := os.Open(*dailyRotoFilePath)
	defer dailyRotoFile.Close()

	if nil != err {
		os.Exit(255)
	}

	dailyRotoParser := dailyRotoCsv.Parser{
		PlayerDeserializer: &dailyrotoSerialization.DefaultPlayerDeserializer{
			IDDeserializer:   &serialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	dailyRotoProjections, err := dailyRotoParser.Deserialize(bufio.NewReader(dailyRotoFile))
	if nil != err {
		os.Exit(255)
	}

	rotogrindersFile, err := os.Open(*rotogrindersFilePath)
	defer rotogrindersFile.Close()

	if nil != err {
		os.Exit(255)
	}

	rotogrindersParser := rotogrindersCsv.Parser{
		PlayerDeserializer: &rotogrindersSerialization.DefaultPlayerDeserializer{
			IDDeserializer:   &serialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	rotogrindersProjections, err := rotogrindersParser.Deserialize(bufio.NewReader(rotogrindersFile))
	if nil != err {
		os.Exit(255)
	}

	sabersimFile, err := os.Open(*sabersimFilePath)
	defer sabersimFile.Close()

	if nil != err {
		os.Exit(255)
	}

	saberSimParser := saberSimCsv.Parser{
		PlayerDeserializer: &saberSimSerialization.DefaultPlayerDeserializer{
			IDDeserializer:   &serialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	saberSimProjections, err := saberSimParser.Deserialize(bufio.NewReader(sabersimFile))
	if nil != err {
		os.Exit(255)
	}

	fmt.Println(awesomeoProjections)
	fmt.Println(dailyRotoProjections)
	fmt.Println(rotogrindersProjections)
	fmt.Println(saberSimProjections)
}
