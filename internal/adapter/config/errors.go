package config

import "errors"

var (
	ErrReadConfigSource = errors.New("can`t read config source")
	ErrUnmarshalYAML    = errors.New("can`t unmarshal yaml")
)
