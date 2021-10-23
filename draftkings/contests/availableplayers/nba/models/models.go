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
	ATLANTAHAWKS     Team = "Atlanta Hawks"
	PHILADELPHA76ERS      = "Philadelphia 76ers"
)

// NBAContestPlayerDetails represents details about a player that is eligible for an NBA contest
type NBAContestPlayerDetails struct {
	Team                   *Team
	Opponent               *Team
	StartTime              *time.Time
	Salary                 float64
	EligibilityByPositions map[ContestPosition]bool
}
