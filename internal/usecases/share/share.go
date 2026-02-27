package share

import (
	"errors"
	"fmt"

	"github.com/OlegLaban/sing_token/internal/domain"
)

type transfer interface {
	Send(data []byte) error
}

type Logger interface {
	Debug(string)
	Error(string, error)
}

type share struct {
	t transfer
	l Logger
}

func (s *share) PutKey(key string) error {
	err := s.t.Send(fmt.Appendf(nil, "KEY%sKEY", key))
	if err != nil {
		s.l.Error("can`t share key ", err)
		return errors.Join(domain.ErrShareKey, err)
	}
	s.l.Debug("key was shared success " + key)
	return nil
}
