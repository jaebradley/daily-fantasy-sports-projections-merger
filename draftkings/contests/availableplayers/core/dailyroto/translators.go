package dailyroto

import "sync"

var once sync.Once

var playerNameByDailyRotoPlayerName = map[string]string{
	"Théo Maledon":          "Theo Maledon",
	"Aleksej Pokuševski":    "Aleksej Pokusevski",
	"C.J. McCollum":         "CJ McCollum",
	"DeAndre Bembry":        "DeAndre' Bembry",
	"DeAndre Hunter":        "De'Andre Hunter",
	"Robert Williams III":   "Robert Williams",
	"Devonte Graham":        "Devonte' Graham",
	"Sviatoslav Mykhailiuk": "Svi Mykhailiuk",
	"DAngelo Russell":       "D'Angelo Russell",
	"Xavier Tillman Sr.":    "Xavier Tillman",
	"Maurice Harkless":      "Moe Harkless",
	"DeAnthony Melton":      "De'Anthony Melton",
	"DeAaron Fox":           "De'Aaron Fox",
}

var defaultPlayerNameTranslator *PlayerNameTranslator

type PlayerNameTranslator struct {
	PlayerNameByDailyRotoPlayerName map[string]string
}

func (t *PlayerNameTranslator) Translate(value string) string {
	name, exist := t.PlayerNameByDailyRotoPlayerName[value]
	if exist {
		return name
	}
	return value
}

func GetDefaultPlayerNameTranslator() *PlayerNameTranslator {
	once.Do(func() {
		defaultPlayerNameTranslator = &PlayerNameTranslator{
			PlayerNameByDailyRotoPlayerName: playerNameByDailyRotoPlayerName,
		}
	})
	return defaultPlayerNameTranslator
}
