package crypter

import "errors"

var (
	ErrBlockCreate    = errors.New("can`t create block")
	ErrTextToShort    = errors.New("ciphertext too short")
	ErrCantReadRandom = errors.New("can`t read random bytes")
	ErrInvalidBase64  = errors.New("invalid base64 data")
)
