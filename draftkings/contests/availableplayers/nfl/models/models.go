package models

import (
	"time"
)

// ContestPosition represents all possible NFL contest positions
type ContestPosition string

const (
	// QUARTERBACK represents the Quarterback position
	QUARTERBACK            ContestPosition = "QUARTERBACK"
	RUNNINGBACK                            = "RUNNING_BACK"
	WIDERECEIVER                           = "WIDE_RECEIVER"
	FLEX                                   = "FLEX"
	TIGHTEND                               = "TIGHT_END"
	DEFENSEANDSPECIALTEAMS                 = "DEFENSE_AND_SPECIAL_TEAMS"
)

// ContestPositions represents a mapping of contest positions to an arbitrary index
var ContestPositions = map[ContestPosition]int{
	QUARTERBACK:            0,
	RUNNINGBACK:            1,
	WIDERECEIVER:           2,
	FLEX:                   3,
	TIGHTEND:               4,
	DEFENSEANDSPECIALTEAMS: 5,
}

type Team string

const (
	ARIZONACARDINALS       Team = "Arizona Cardinals"
	ATLANTAFALCONS              = "Atlanta Falcons"
	BALTIMORERAVENS             = "Baltimore Ravens"
	BUFFALOBILLS                = "Buffalo Bills"
	CAROLINAPANTHERS            = "Carolina Panthers"
	CHICAGOBEARS                = "Chicago Bears"
	CINCINNATIBENGALS           = "Cincinnati Bengals"
	CLEVELANDBROWNS             = "Cleveland Browns"
	DALLASCOWBOYS               = "Dallas Cowboys"
	DENVERBRONCOS               = "Denver Broncos"
	DETROITLIONS                = "Detroit Lions"
	GREENBAYPACKERS             = "Green Bay Packers"
	HOUSTONTEXANS               = "Houston Texans"
	INDIANAPOLISCOLTS           = "Indianapolis Colts"
	JACKSONVILLEJAGUARS         = "Jacksonville Jaguars"
	KANSASCITYCHIEFS            = "Kansas City Chiefs"
	LASVEGASRAIDERS             = "Las Vegas Raiders"
	LOSANGELESCHARGERS          = "Los Angeles Chargers"
	LOSANGELESRAMS              = "Los Angeles Rams"
	MIAMIDOLPHINS               = "Miami Dolphins"
	MINNESOTAVIKINGS            = "Minnesota Vikings"
	NEWENGLANDPATRIOTS          = "New England Patriots"
	NEWORLEANSSAINTS            = "New Orleans Saints"
	NEWYORKGIANTS               = "New York Giants"
	NEWYORKJETS                 = "New York Jets"
	PHILADELPHIAEAGLES          = "Philadelphia Eagles"
	PITTSBURGHSTEELERS          = "Pittsburgh Steelers"
	SANFRANCISCO49ERS           = "San Francisco 49ers"
	SEATTLESEAHAWKS             = "Seattle Seahawks"
	TAMPABAYBUCCANEERS          = "Tampa Bay Buccaneers"
	TENNESSEETITANS             = "Tennessee Titans"
	WASHINGTONFOOTBALLTEAM      = "Washington Football Team"
)

// Player represents a Player in DraftKings
type Player struct {
	Id   int
	Name string
}

// NFLContestPlayerDetails represents details about a player that is eligible for an NFL contest
type NFLContestPlayerDetails struct {
	Team                   *Team
	Opponent               *Team
	StartTime              *time.Time
	Salary                 float32
	EligibilityByPositions map[ContestPosition]bool
}
