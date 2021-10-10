package models

import (
  "time"
)

// ContestPosition represents all possible NFL contest positions
type ContestPosition string

const(
    // QUARTERBACK represents the Quarterback position
    QUARTERBACK ContestPosition = "QUARTERBACK"
    RUNNINGBACK = "RUNNING_BACK"
    WIDERECEIVER = "WIDE_RECEIVER"
    FLEX = "FLEX"
    TIGHTEND = "TIGHT_END"
    DEFENSEANDSPECIALTEAMS = "DEFENSE_AND_SPECIAL_TEAMS"
)

// Player represents a Player in DraftKings
type Player struct {
  id  int
  name string
}

// NFLContestPlayerDetails represents details about a player that is eligible for an NFL contest
type NFLContestPlayerDetails struct {
  team Team
  opponent Team
  startTime time.Time
  salary float
  eligibilityByPositions map[ContestPosition]bool
}
