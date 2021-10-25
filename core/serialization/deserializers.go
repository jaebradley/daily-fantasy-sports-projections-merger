package serialization

import "strconv"

type PlayerIDDeserializer interface {
	Deserialize(value string) (int64, error)
}

type PlayerNameDeserializer interface {
	Deserialize(value string) (string, error)
}

type ProjectionDeserializer interface {
	Deserialize(value string) (float64, error)
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
	if "" == value {
		return 0, nil
	}
	return strconv.ParseFloat(value, 64)
}
