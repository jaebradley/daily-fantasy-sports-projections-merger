package csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"

	"github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/models"
	"github.com/jaebradley/daily-fantasy-sports-projections-merger/sabersim/serialization"
)

// Deserializer deserializes a File into a mapping of player to the player's contest details
type Deserializer interface {
	Deserialize(bytes *bytes.Buffer) (map[models.Player]float64, error)
}

type Parser struct {
	PlayerDeserializer     serialization.PlayerDeserializer
	ProjectionDeserializer serialization.ProjectionDeserializer
}

func (p *Parser) Deserialize(bytes *bytes.Buffer) (map[models.Player]float64, error) {
	projectionsByPlayer := make(map[models.Player]float64)
	reader := csv.NewReader(bytes)

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
		player, err := p.PlayerDeserializer.Deserialize(record[0], record[1])
		_, exist := projectionsByPlayer[*player]
		if exist {
			return nil, errors.New("duplicate player found")
		}
		projection, err := p.ProjectionDeserializer.Deserialize(record[8])
		if nil != err {
			return nil, err
		}
		projectionsByPlayer[*player] = projection
	}
	return projectionsByPlayer, nil
}
