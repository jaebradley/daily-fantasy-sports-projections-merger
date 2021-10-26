package models

import (
	draftKingsNbaModels "github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/nba/models"
)

type Site string

const (
	AWESOMEO     Site = "Awesomeo"
	DAILYROTO         = "DailyRoto"
	ROTOGRINDERS      = "RotoGrinders"
	SABERSIM          = "SaberSim"
)

type NBAProjectionDetails struct {
	PlayerDetails     draftKingsNbaModels.NBAContestPlayerDetails
	ProjectionsBySite map[Site]float64
}
