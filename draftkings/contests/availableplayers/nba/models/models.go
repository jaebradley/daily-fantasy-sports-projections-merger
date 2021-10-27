package models

import "time"

type ContestPosition string

const (
	POINTGUARD    ContestPosition = "Point Guard"
	SHOOTINGGUARD                 = "Shooting Guard"
	SMALLFORWARD                  = "Small Forward"
	POWERFORWARD                  = "Power Forward"
	CENTER                        = "Center"
	GUARD                         = "Guard"
	FORWARD                       = "Forward"
	UTILITY                       = "Utility"
)

// ContestPositions represents a mapping of contest positions to an arbitrary index
var ContestPositions = map[ContestPosition]int{
	POINTGUARD:    0,
	SHOOTINGGUARD: 1,
	SMALLFORWARD:  2,
	POWERFORWARD:  3,
	CENTER:        4,
	GUARD:         5,
	FORWARD:       6,
	UTILITY:       7,
}

type Team string

const (
	ATLANTAHAWKS          Team = "Atlanta Hawks"
	BOSTONCELTICS              = "Boston Celtics"
	BROOKLYNNETS               = "Brooklyn Nets"
	CHARLOTTEHORNETS           = "Charlotte Hornets"
	CHICAGOBULLS               = "Chicago Bulls"
	CLEVELANDCAVALIERS         = "Cleveland Cavaliers"
	DALLASMAVERICKS            = "Dallas Mavericks"
	DENVERNUGGETS              = "Denver Nuggets"
	DETROITPISTONS             = "Detroit Pistons"
	GOLDENSTATEWARRIORS        = "Golden State Warriors"
	HOUSTONROCKETS             = "Houston Rockets"
	INDIANAPACERS              = "Indiana Pacers"
	LOSANGELESCLIPPERS         = "Los Angeles Clippers"
	LOSANGELESLAKERS           = "Los Angeles Lakers"
	MEMPHISGRIZZLIES           = "Memphis Grizzlies"
	MIAMIHEAT                  = "Miami Heat"
	MILWAUKEEBUCKS             = "Milwaukee Bucks"
	MINNESOTATIMBERWOLVES      = "Minnesota Timberwolves"
	NEWORLEANSPELICANS         = "New Orleans Pelicans"
	NEWYORKKNICKS              = "New York Knicks"
	OKLAHOMACITYTHUNDER        = "Oklahoma City Thunder"
	ORLANDOMAGIC               = "Orlando Magic"
	PHILADELPHA76ERS           = "Philadelphia 76ers"
	PHOENIXSUNS                = "Phoenix Suns"
	PORTLANDTRAILBLAZERS       = "Portland Trail Blazers"
	SACRAMENTOKINGS            = "Sacramento Kings"
	SANANTONIOSPURS            = "San Antonio Spurs"
	TORONTORAPTORS             = "Toronto Raptors"
	UTAHJAZZ                   = "Utah Jazz"
	WASHINGTONWIZARDS          = "Washington Wizards"
)

// NBAContestPlayerDetails represents details about a player that is eligible for an NBA contest
type NBAContestPlayerDetails struct {
	Team                   *Team
	Opponent               *Team
	StartTime              *time.Time
	Salary                 float64
	EligibilityByPositions map[ContestPosition]bool
}
