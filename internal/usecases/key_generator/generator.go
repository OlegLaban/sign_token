package keygenerator

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/OlegLaban/sing_token/internal/domain"
)

type Logger interface {
	Infof(string, ...any)
	Debugf(string, ...any)
	Error(string, error)
}

type Crypter interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type generator struct {
	c Crypter
	l Logger
}

type key struct {
	Indentify string `json:"id"`
	DateTime  string `json:"date_time"`
}

func NewGenerator(c Crypter, l Logger) *generator {
	return &generator{c: c, l: l}
}

func (g *generator) Generate(p domain.Payload) (string, error) {
	tokenData, err := g.convertPayload(p)
	if err != nil {
		g.l.Error("can`t encode payload to string %s", err)
		return "", errors.Join(ErrCantConvertKey, err)
	}
	g.l.Debugf("payload was converted to tokenData - %s", tokenData)
	encryptedText, err := g.c.Encrypt(tokenData)
	if err != nil {
		g.l.Error("can`t encrypt payload", err)
		return "", errors.Join(ErrCantEncryptKey, err)
	}
	g.l.Debugf("token data was encrypted - %s", encryptedText)

	return encryptedText, nil
}

func (g *generator) convertPayload(p domain.Payload) (string, error) {
	keyStruct := key{
		Indentify: p.FullName,
		DateTime:  strconv.Itoa(p.Date),
	}

	data, err := json.Marshal(keyStruct)
	if err != nil {
		g.l.Error("can`t marshal key struct to json", err)
		return "", err
	}
	g.l.Debugf("payload was converted to json - %d", len(data))

	return string(data), nil
}
