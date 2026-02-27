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
	// if err != nil {
	// 	s.l.Error("can`t share key ", err)
	// 	return errors.Join(domain.ErrShareKey, err)
	// }
	s.l.Debug("key was shared success " + key)
	return nil
}
