package csv

import (
	"encoding/csv"
	"errors"
	"io"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/core/serialization"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/models"
)

// Deserializer deserializes a File into a mapping of player to the player's contest details
type Deserializer interface {
	Deserialize(r io.Reader) (map[models.Player]float64, error)
}

type Parser struct {
	PlayerNameDeserializer serialization.PlayerNameDeserializer
	ProjectionDeserializer serialization.ProjectionDeserializer
}

func (p *Parser) Deserialize(r io.Reader) (map[string]float64, error) {
	projectionsByPlayer := make(map[string]float64)
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
		playerName, err := p.PlayerNameDeserializer.Deserialize(record[0])
		if nil != err {
			return nil, err
		}
		_, exist := projectionsByPlayer[playerName]
		if exist {
			return nil, errors.New("duplicate player found")
		}
		projection, err := p.ProjectionDeserializer.Deserialize(record[1])
		if nil != err {
			return nil, err
		}
		projectionsByPlayer[playerName] = projection
	}
	return projectionsByPlayer, nil
}
