package core

type Translator interface {
	Translate(value string) string
}
