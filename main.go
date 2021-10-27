package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	awesomeoCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/awesomeo/csv"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	dailyRotoCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/csv"
	dailyrotoSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/serialization"
	draftKingsCoreModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/models"
	draftKingsCoreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/serialization"
	draftKingsNbaCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/csv"
	draftKingsNbaModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
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
	draftKingsFilePath := flag.String("draftkings", "", "path to DraftKings available players as a CSV file")
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

	draftKingsFile, err := os.Open(*draftKingsFilePath)
	defer draftKingsFile.Close()

	if nil != err {
		os.Exit(255)
	}

	draftKingsParser := draftKingsNbaCsv.Parser{
		TimeDeserializer:   &draftKingsCoreSerialization.DefaultContestStartTimeDeserializer{},
		SalaryDeserializer: &draftKingsCoreSerialization.DefaultSalaryDeserializer{},
		PlayerDeserializer: &draftKingsCoreSerialization.DefaultPlayerDeserializer{
			IdDeserializer:   &serialization.DefaultPlayerIDDeserializer{},
			NameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		},
		ContestPositionsDeserializer: &draftKingsNbaCsv.ContestPositionsDeserializer{
			PositionDeserializer: &draftKingsNbaCsv.ContestPositionDeserializer{
				PositionsByAbbreviation: map[string]draftKingsNbaModels.ContestPosition{
					"PG":   draftKingsNbaModels.POINTGUARD,
					"SG":   draftKingsNbaModels.SHOOTINGGUARD,
					"SF":   draftKingsNbaModels.SMALLFORWARD,
					"PF":   draftKingsNbaModels.POWERFORWARD,
					"C":    draftKingsNbaModels.CENTER,
					"G":    draftKingsNbaModels.GUARD,
					"F":    draftKingsNbaModels.FORWARD,
					"UTIL": draftKingsNbaModels.UTILITY,
				},
			},
			Separator: '/',
		},
		TeamDeserializer: &draftKingsNbaCsv.TeamAbbreviationDeserializer{
			TeamsByAbbreviation: map[string]draftKingsNbaModels.Team{
				"ATL": draftKingsNbaModels.ATLANTAHAWKS,
				"BOS": draftKingsNbaModels.BOSTONCELTICS,
				"CHA": draftKingsNbaModels.CHARLOTTEHORNETS,
				"CHI": draftKingsNbaModels.CHICAGOBULLS,
				"CLE": draftKingsNbaModels.CLEVELANDCAVALIERS,
				"DAL": draftKingsNbaModels.DALLASMAVERICKS,
				"DEN": draftKingsNbaModels.DENVERNUGGETS,
				"DET": draftKingsNbaModels.DETROITPISTONS,
				"GSW": draftKingsNbaModels.GOLDENSTATEWARRIORS,
				"HOU": draftKingsNbaModels.HOUSTONROCKETS,
				"IND": draftKingsNbaModels.INDIANAPACERS,
				"LAC": draftKingsNbaModels.LOSANGELESCLIPPERS,
				"LAL": draftKingsNbaModels.LOSANGELESCLIPPERS,
				"MEM": draftKingsNbaModels.MEMPHISGRIZZLIES,
				"MIA": draftKingsNbaModels.MIAMIHEAT,
				"MIL": draftKingsNbaModels.MILWAUKEEBUCKS,
				"MIN": draftKingsNbaModels.MINNESOTATIMBERWOLVES,
				"NOP": draftKingsNbaModels.NEWORLEANSPELICANS,
				"NYK": draftKingsNbaModels.NEWYORKKNICKS,
				"OKC": draftKingsNbaModels.OKLAHOMACITYTHUNDER,
				"ORL": draftKingsNbaModels.ORLANDOMAGIC,
				"PHI": draftKingsNbaModels.PHILADELPHA76ERS,
				"PHX": draftKingsNbaModels.PHOENIXSUNS,
				"POR": draftKingsNbaModels.PORTLANDTRAILBLAZERS,
				"SAC": draftKingsNbaModels.SACRAMENTOKINGS,
				"SAS": draftKingsNbaModels.SANANTONIOSPURS,
				"TOR": draftKingsNbaModels.TORONTORAPTORS,
				"UTA": draftKingsNbaModels.UTAHJAZZ,
				"WAS": draftKingsNbaModels.WASHINGTONWIZARDS,
			},
		},
		OpponentDeserializer: &draftKingsNbaCsv.TeamAbbreviationDeserializer{
			TeamsByAbbreviation: map[string]draftKingsNbaModels.Team{
				"PHI": draftKingsNbaModels.PHILADELPHA76ERS,
			},
		},
	}

	availableDraftKingsPlayers, err := draftKingsParser.Deserialize(bufio.NewReader(draftKingsFile))
	if nil != err {
		fmt.Println(err)
		os.Exit(255)
	}

	dailyRotoProjectionsByPlayerName := make(map[string]float64)

	for dailyRotoPlayer, projection := range dailyRotoProjections {
		dailyRotoProjectionsByPlayerName[dailyRotoPlayer.Name] = projection
	}

	rotogrindersProjectionsByPlayerName := make(map[string]float64)

	for rotogrindersPlayer, projection := range rotogrindersProjections {
		rotogrindersProjectionsByPlayerName[rotogrindersPlayer.Name] = projection
	}

	sabersimProjectionsByPlayerName := make(map[string]float64)

	for sabersimPlayer, projection := range saberSimProjections {
		sabersimProjectionsByPlayerName[sabersimPlayer.Name] = projection
	}

	projectionsByPlayer := make(map[draftKingsCoreModels.Player]models.NBAProjectionDetails)

	for player, details := range availableDraftKingsPlayers {
		projectionsBySite := make(map[models.Site]float64)
		awesomeoProjection, exist := awesomeoProjections[player.Name]
		if exist {
			projectionsBySite[models.AWESOMEO] = awesomeoProjection
		}

		dailyRotoProjection, exist := dailyRotoProjectionsByPlayerName[player.Name]
		if exist {
			projectionsBySite[models.DAILYROTO] = dailyRotoProjection
		}

		rotogrindersProjection, exist := rotogrindersProjectionsByPlayerName[player.Name]
		if exist {
			projectionsBySite[models.ROTOGRINDERS] = rotogrindersProjection
		}

		sabersimProjection, exist := sabersimProjectionsByPlayerName[player.Name]
		if exist {
			projectionsBySite[models.SABERSIM] = sabersimProjection
		}

		projectionsByPlayer[player] = models.NBAProjectionDetails{
			PlayerDetails:     details,
			ProjectionsBySite: projectionsBySite,
		}
	}

	file, err := os.Create("combined.csv")
	if nil != err {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for player, details := range projectionsByPlayer {
		r := make([]string, 0, 7)
		r = append(r, strconv.FormatInt(player.Id, 10))
		r = append(r, player.Name)
		r = append(r, strconv.FormatFloat(details.PlayerDetails.Salary, 'f', -1, 64))

		awesomeoProjection, exist := details.ProjectionsBySite[models.AWESOMEO]
		if false == exist {
			awesomeoProjection = 0
		}
		r = append(r, strconv.FormatFloat(awesomeoProjection, 'f', -1, 64))

		dailyRotoProjection, exist := details.ProjectionsBySite[models.DAILYROTO]
		if false == exist {
			dailyRotoProjection = 0
		}
		r = append(r, strconv.FormatFloat(dailyRotoProjection, 'f', -1, 64))

		rotogrindersProjection, exist := details.ProjectionsBySite[models.ROTOGRINDERS]
		if false == exist {
			rotogrindersProjection = 0
		}
		r = append(r, strconv.FormatFloat(rotogrindersProjection, 'f', -1, 64))

		sabersimProjection, exist := details.ProjectionsBySite[models.SABERSIM]
		if false == exist {
			sabersimProjection = 0
		}
		r = append(r, strconv.FormatFloat(sabersimProjection, 'f', -1, 64))
		err := writer.Write(r)
		if nil != err {
			fmt.Println(err)
		}
	}
}
