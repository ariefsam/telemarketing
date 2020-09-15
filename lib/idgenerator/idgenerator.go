package idgenerator

import (
	"github.com/teris-io/shortid"
)

type IDGenerator struct{}

func (i *IDGenerator) Generate() (id string) {
	id, _ = shortid.Generate()
	return
}
