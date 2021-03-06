package models

import (
	draftKingsNbaModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
)

type Site string

const (
	AWESOMEO     Site = "Awesomeo"
	DAILYROTO         = "DailyRoto"
	ENTERTHERUN       = "EnterTheRun"
	ROTOGRINDERS      = "RotoGrinders"
	SABERSIM          = "SaberSim"
)

var Sites = map[Site]int{
	AWESOMEO:     0,
	DAILYROTO:    1,
	ENTERTHERUN:  2,
	ROTOGRINDERS: 3,
	SABERSIM:     4,
}

type NBAProjectionDetails struct {
	PlayerDetails     draftKingsNbaModels.NBAContestPlayerDetails
	ProjectionsBySite map[Site]float64
}
