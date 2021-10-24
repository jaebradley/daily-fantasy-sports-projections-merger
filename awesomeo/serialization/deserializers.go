package serialization

import (
	coreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/dailyroto/models"
)

type PlayerDeserializer interface {
	Deserialize(id string, name string) (*models.Player, error)
}

type DefaultPlayerDeserializer struct {
	IDDeserializer   coreSerialization.PlayerIDDeserializer
	NameDeserializer coreSerialization.PlayerNameDeserializer
}

func (d *DefaultPlayerDeserializer) Deserialize(id string, name string) (*models.Player, error) {
	playerID, err := d.IDDeserializer.Deserialize(id)
	if nil != err {
		return nil, err
	}

	playerName, err := d.NameDeserializer.Deserialize(name)
	if nil != err {
		return nil, err
	}

	return &models.Player{
		ID:   playerID,
		Name: playerName,
	}, nil
}
