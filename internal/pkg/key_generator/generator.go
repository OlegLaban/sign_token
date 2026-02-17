package keygenerator

import (
	"github.com/OlegLaban/sing_token/internal/domain"
)

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

func (g *generator) Generate(p domain.Payload) (string, error) {
	return "", nil
}
