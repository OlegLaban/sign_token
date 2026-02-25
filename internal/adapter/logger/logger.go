package logger

import (
	"log"
)

type logger struct{}

func New() *logger {
	return &logger{}
}

func (l *logger) Infof(s string, arg ...any) {
	log.Default().Printf("[INFO]"+s+"\n", arg...)
}

func (l *logger) Debugf(s string, arg ...any) {
	log.Default().Printf("[DEBUG] "+s+"\n", arg...)
}

func (l *logger) Error(s string, err error) {
	log.Default().Print("[ERROR]"+s+"\n", err.Error())
}
