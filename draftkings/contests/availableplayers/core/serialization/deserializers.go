package serialization

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/draftkings/contests/availableplayers/core/models"
)

// PlayerIDDeserializer deserializes a string into a Player Id
type PlayerIDDeserializer interface {
	Deserialize(value string) (int64, error)
}

// PlayerNameDeserializer deserializes a string into a Player Name
type PlayerNameDeserializer interface {
	Deserialize(value string) (string, error)
}

// PlayerDeserializer deserializes an id string and a name string into a Player
type PlayerDeserializer interface {
	Deserialize(id string, name string) (*models.Player, error)
}

// SalaryDeserializer deserializes a string into a salary
type SalaryDeserializer interface {
	Deserialize(value string) (float64, error)
}

// TimeDeserializer deserializes a string into a time.Time
type TimeDeserializer interface {
	Deserialize(value string) (*time.Time, error)
}

type DefaultSalaryDeserializer struct {
}

func (d *DefaultSalaryDeserializer) Deserialize(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

type DefaultPlayerIDDeserializer struct {
}

func (d *DefaultPlayerIDDeserializer) Deserialize(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

type DefaultPlayerNameDeserializer struct {
}

func (d *DefaultPlayerNameDeserializer) Deserialize(value string) (string, error) {
	return value, nil
}

type DefaultPlayerDeserializer struct {
	IdDeserializer   PlayerIDDeserializer
	NameDeserializer PlayerNameDeserializer
}

func (d *DefaultPlayerDeserializer) Deserialize(id string, name string) (*models.Player, error) {
	playerID, err := d.IdDeserializer.Deserialize(id)
	if nil != err {
		return nil, err
	}

	playerName, err := d.NameDeserializer.Deserialize(name)
	if nil != err {
		return nil, err
	}

	return &models.Player{
		Id:   playerID,
		Name: playerName,
	}, nil
}

// DefaultContestStartTimeDeserializer deserializer contest start times
type DefaultContestStartTimeDeserializer struct {
}

// Deserialize deserializes a string into a time
func (d *DefaultContestStartTimeDeserializer) Deserialize(value string) (*time.Time, error) {
	parts := strings.Split(value, " ")
	if 4 != len(parts) {
		return nil, errors.New("expected exactly 4 parts")
	}

	if "ET" != parts[3] {
		return nil, errors.New("unexpected time zone")
	}

	easternTimeZone, err := time.LoadLocation("America/New_York")
	if nil != err {
		return nil, errors.New("unable to load Eastern Time Zone")
	}

	parsedTime, err := time.ParseInLocation("01/02/2006 03:04PM", strings.Join(parts[1:3], " "), easternTimeZone)
	if nil != err {
		return nil, err
	}
	return &parsedTime, nil
}
