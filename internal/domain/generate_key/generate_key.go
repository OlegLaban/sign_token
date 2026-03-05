package generatekey

import (
	"io"
	"time"

	"github.com/OlegLaban/sing_token/internal/domain"
)

type Input interface {
	Get() string
	Set(string)
}

type Logger interface {
	Error(string, error)
	Debug(string)
}

type Crypt interface {
	Generate(domain.Payload) (string, error)
}

type Share interface {
	Send(string) error
}

type generator struct {
	out io.Writer
	i   Input
	l   Logger
	c   Crypt
	s   Share
}

func New(out io.Writer, i Input, l Logger, c Crypt, s Share) *generator {
	return &generator{out: out, i: i, l: l, c: c, s: s}
}

func (g *generator) Make() func() {
	return func() {
		cryptKey, err := g.c.Generate(domain.NewPayload(g.i.Get(), int(time.Now().Unix())))
		if err != nil {
			g.l.Error("can`t generate key", err)
			return
		}
		g.l.Debug("key was generated success")
		err = g.s.Send(cryptKey)
		if err != nil {
			g.l.Error("can`t share key\n", err)
			return
		}
		g.l.Debug("key was shared successfuly")
		g.out.Write([]byte(cryptKey))
		g.i.Set("")
		g.l.Debug("key generated logic was success")
	}
}
