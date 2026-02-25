package config

import (
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Crypto Crypt `yaml:"crypto"`
}

type Crypt struct {
	Key string `yaml:"key"`
}

func Parse(r io.Reader) (Config, error) {
	var c Config
	data, err := io.ReadAll(r)
	if err != nil {
		return Config{}, errors.Join(ErrReadConfigSource, err)
	}
	if err = yaml.Unmarshal(data, &c); err != nil {
		return Config{}, errors.Join(ErrUnmarshalYAML, err)
	}
	return c, nil
}
