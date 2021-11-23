package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func (c *Config) InitConfig() *Config {
	file, err := ioutil.ReadFile("files/config-covid19-development.yml")
	if err != nil {
		log.Fatal("[config][InitConfig]", err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatal("[config][InitConfig]", err)
	}

	return c
}
