package keygenerator

import "errors"

var (
	ErrCantConvertKey = errors.New("can`t convert payload to key")
	ErrCantEncryptKey = errors.New("can`t encrypt key")
)
