package share

import (
	"fmt"
)

type transfer interface {
	Send(data []byte)
}

type Logger interface {
	Debug(string)
	Error(string, error)
}

type share struct {
	t transfer
	l Logger
}

func New(t transfer, l Logger) *share {
	return &share{t: t, l: l}
}

func (s *share) PutKey(key string) error {
	s.t.Send(fmt.Appendf(nil, "KEY%sKEY", key))
	s.l.Debug("key was shared success " + key)
	return nil
}
