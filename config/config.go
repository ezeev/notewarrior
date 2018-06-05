package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	MarkDownPath string `yaml:"mdpath"`
}

func NewConfig(path string) (*Config, error) {
	c := Config{}
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, &c)
	return &c, err
}
