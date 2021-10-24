package serialization

import (
	"strconv"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/models"
)

type PlayerIDDeserializer interface {
	Deserialize(value string) (int64, error)
}

type PlayerNameDeserializer interface {
	Deserialize(value string) (string, error)
}

type ProjectionDeserializer interface {
	Deserialize(value string) (float64, error)
}

type PlayerDeserializer interface {
	Deserialize(id string, name string) (*models.Player, error)
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

type DefaultProjectionDeserializer struct {
}

func (d *DefaultProjectionDeserializer) Deserialize(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

type DefaultPlayerDeserializer struct {
	IdDeserializer   PlayerIDDeserializer
	NameDeserializer PlayerNameDeserializer
}

func (d *DefaultPlayerDeserializer) Deserialize(id string, name string) (*models.Player, error) {
	playerId, err := d.IdDeserializer.Deserialize(id)
	if nil != err {
		return nil, err
	}

	playerName, err := d.NameDeserializer.Deserialize(name)
	if nil != err {
		return nil, err
	}

	return &models.Player{
		Id:   playerId,
		Name: playerName,
	}, nil
}
