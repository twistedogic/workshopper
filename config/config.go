package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func merge(configs ...Config) Config {
	c := Config{}
	for _, cfg := range configs {
		c.Stages = append(c.Stages, cfg.Stages...)
		c.Exercises = append(c.Exercises, cfg.Exercises...)
	}
	return c
}

type Config struct {
	dir      string   `yaml:"-"`
	Includes []string `yaml:"include"`
	Stages   []string `yaml:"stages"`
	Exercises
}

func (c Config) Validate() error { return nil }

func (c Config) Merge(o Config) Config {
	return merge(c, o)
}

func (c Config) Expand() (Config, error) {
	cfg := Config{}
	for _, f := range c.Includes {
		nested, err := FromYAMLFile(filepath.Join(c.dir, f))
		if err != nil {
			return cfg, err
		}
		cfg = cfg.Merge(nested)
	}
	return cfg, nil
}

func FromYAML(b []byte) (Config, error) {
	var c Config
	return c, yaml.Unmarshal(b, &c)
}

func FromYAMLFile(filename string) (Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	cfg, err := FromYAML(b)
	if err != nil {
		return Config{}, err
	}
	cfg.dir = filepath.Dir(filename)
	return cfg.Expand()
}
