package linux

import (
	"log"
	"time"
)

type linuxAPP struct{}

func NewAPP() *linuxAPP {
	return &linuxAPP{}
}

func (la *linuxAPP) Run() {
	now := time.Now().Unix()
	log.Println(now)
}
