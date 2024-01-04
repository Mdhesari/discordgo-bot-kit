package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(path string) Config {
	var k = koanf.New(".")

	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		panic(err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return cfg
}
