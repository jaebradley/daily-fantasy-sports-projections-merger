package csv

import (
	"encoding/csv"
	"errors"
	"io"

	coreSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/models"
	rotogrindersSerialization "github.com/jaebradley/daily-fantasy-sports-projections-merger/rotogrinders/serialization"
)

// Deserializer deserializes a File into a mapping of player to the player's contest details
type Deserializer interface {
	Deserialize(r io.Reader) (map[models.Player]float64, error)
}

type Parser struct {
	PlayerDeserializer     rotogrindersSerialization.PlayerDeserializer
	ProjectionDeserializer coreSerialization.ProjectionDeserializer
}

func (p *Parser) Deserialize(r io.Reader) (map[models.Player]float64, error) {
	projectionsByPlayer := make(map[models.Player]float64)
	reader := csv.NewReader(r)

	_, err := reader.Read()

	if err == io.EOF {
		return projectionsByPlayer, nil
	}

	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			return projectionsByPlayer, nil
		}

		if err != nil {
			return nil, err
		}
		player, err := p.PlayerDeserializer.Deserialize(record[0], record[4])
		_, exist := projectionsByPlayer[*player]
		if exist {
			return nil, errors.New("duplicate player found")
		}
		projection, err := p.ProjectionDeserializer.Deserialize(record[5])
		if nil != err {
			return nil, err
		}
		projectionsByPlayer[*player] = projection
	}
	return projectionsByPlayer, nil
}
