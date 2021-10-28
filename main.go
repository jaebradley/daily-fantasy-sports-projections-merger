package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"

	awesomeoCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/awesomeo/csv"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	dailyRotoCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/csv"
	dailyRotoModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/models"
	dailyrotoSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/serialization"
	coreDailyRoto "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/dailyroto"
	draftKingsCoreModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/models"
	draftKingsCoreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/serialization"
	draftKingsNbaCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/csv"
	draftKingsNbaModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
	etrCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/entertherun/csv"
	rotogrindersCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/csv"
	rotoGrindersModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/models"
	rotogrindersSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/serialization"
	saberSimCsv "github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/csv"
	saberSimModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/models"
	saberSimSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/serialization"
)

func main() {
	awesomeoFilePath := flag.String("awesomeo", "", "path to Awesomeo projections downloaded as a CSV file")
	dailyRotoFilePath := flag.String("dailyroto", "", "path to DailyRoto projections downloaded as a CSV file")
	rotogrindersFilePath := flag.String("rotogrinders", "", "path to RotoGrinders projections downloaded as a CSV file")
	sabersimFilePath := flag.String("sabersim", "", "path to Sabersim projections downloaded as a CSV file")
	draftKingsFilePath := flag.String("draftkings", "", "path to DraftKings available players as a CSV file")
	etrFilePath := flag.String("etr", "", "path to EnterTheRun projections downloaded as a CSV file")
	flag.Parse()

	draftKingsFile, err := os.Open(*draftKingsFilePath)
	defer draftKingsFile.Close()

	if nil != err {
		log.Fatal(err)
	}

	awesomeoFile, err := os.Open(*awesomeoFilePath)
	defer awesomeoFile.Close()

	if nil != err {
		log.Fatal(err)
	}

	dailyRotoFile, err := os.Open(*dailyRotoFilePath)
	defer dailyRotoFile.Close()

	if nil != err {
		log.Fatal(err)
	}

	rotogrindersFile, err := os.Open(*rotogrindersFilePath)
	defer rotogrindersFile.Close()

	if nil != err {
		log.Fatal(err)
	}

	sabersimFile, err := os.Open(*sabersimFilePath)
	defer sabersimFile.Close()

	if nil != err {
		log.Fatal(err)
	}

	etrFile, err := os.Open(*etrFilePath)
	defer etrFile.Close()

	if nil != err {
		log.Fatal(err)
	}

	awesomeoParser := awesomeoCsv.Parser{
		PlayerNameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	awesomeoProjections, err := awesomeoParser.Deserialize(bufio.NewReader(awesomeoFile))
	if nil != err {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	etrParser := etrCsv.Parser{
		PlayerNameDeserializer: &serialization.DefaultPlayerNameDeserializer{},
		ProjectionDeserializer: &serialization.DefaultProjectionDeserializer{},
	}

	etrProjections, err := etrParser.Deserialize(bufio.NewReader(etrFile))
	if nil != err {
		log.Fatal(err)
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
				"BKN": draftKingsNbaModels.BROOKLYNNETS,
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
		log.Fatal(err)
	}

	playerProjectionsBySite := map[models.Site]map[string]float64{
		models.AWESOMEO:     awesomeoProjections,
		models.DAILYROTO:    makeDailyRotoProjectionsByPlayerName(dailyRotoProjections),
		models.ENTERTHERUN:  etrProjections,
		models.ROTOGRINDERS: makeRotoGrindersProjectionsByPlayerName(rotogrindersProjections),
		models.SABERSIM:     makeSaberSimProjectionsByPlayerName(saberSimProjections),
	}
	projectionsByPlayer := make(map[draftKingsCoreModels.Player]models.NBAProjectionDetails)

	for player, details := range availableDraftKingsPlayers {
		projectionsBySite := make(map[models.Site]float64)

		for site, playerProjections := range playerProjectionsBySite {
			projection, exists := playerProjections[player.Name]
			if exists {
				projectionsBySite[site] = projection
			}
		}
		projectionsByPlayer[player] = models.NBAProjectionDetails{
			PlayerDetails:     details,
			ProjectionsBySite: projectionsBySite,
		}
	}

	file, err := os.Create("combined.csv")
	if nil != err {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"DraftKings Player ID", "Name", "Salary", "Awesomeo", "DailyRoto", "RotoGrinders", "SaberSim", "EnterTheRun"}
	writer.Write(header)

	for player, details := range projectionsByPlayer {
		r := make([]string, 0, 8)
		r = append(r, strconv.FormatInt(player.Id, 10))
		r = append(r, player.Name)
		r = append(r, strconv.FormatFloat(details.PlayerDetails.Salary, 'f', -1, 64))
		r = appendSiteProjection(models.AWESOMEO, details.ProjectionsBySite, r)
		r = appendSiteProjection(models.DAILYROTO, details.ProjectionsBySite, r)
		r = appendSiteProjection(models.ROTOGRINDERS, details.ProjectionsBySite, r)
		r = appendSiteProjection(models.SABERSIM, details.ProjectionsBySite, r)
		r = appendSiteProjection(models.ENTERTHERUN, details.ProjectionsBySite, r)

		err := writer.Write(r)
		if nil != err {
			log.Fatal(err)
		}
	}
}

func makeDailyRotoProjectionsByPlayerName(projectionsByPlayer map[dailyRotoModels.Player]float64) map[string]float64 {
	playerNameTranslator := coreDailyRoto.GetDefaultPlayerNameTranslator()
	projectionsByName := make(map[string]float64)

	for player, projection := range projectionsByPlayer {
		name := playerNameTranslator.Translate(player.Name)
		projectionsByName[name] = projection
	}

	return projectionsByName
}

func makeRotoGrindersProjectionsByPlayerName(projectionsByPlayer map[rotoGrindersModels.Player]float64) map[string]float64 {
	projectionsByName := make(map[string]float64)

	for player, projection := range projectionsByPlayer {
		projectionsByName[player.Name] = projection
	}

	return projectionsByName
}

func makeSaberSimProjectionsByPlayerName(projectionsByPlayer map[saberSimModels.Player]float64) map[string]float64 {
	projectionsByName := make(map[string]float64)

	for player, projection := range projectionsByPlayer {
		projectionsByName[player.Name] = projection
	}

	return projectionsByName
}

func appendSiteProjection(site models.Site, projectionsBySite map[models.Site]float64, record []string) []string {
	projection, exist := projectionsBySite[site]
	if false == exist {
		projection = 0
	}
	return append(record, strconv.FormatFloat(projection, 'f', -1, 64))
}
