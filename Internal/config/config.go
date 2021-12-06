package config

import (
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v2"
)

func (c *Config) InitConfig() *Config {
	file, err := ioutil.ReadFile("files/config-covid19-development.yml")
	if err != nil {
		log.Fatal("[config][InitConfig]", err)
	}

	jsonKey, err := ioutil.ReadFile("files/storage-1bb41-firebase-adminsdk-esal9-7d4133437b.json")
	if err != nil {
		log.Fatal("[config][jsonKey]", err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatal("[config][InitConfig]", err)
	}

	c.Conf, _ = google.JWTConfigFromJSON(jsonKey)
	c.Server.BucketName = "storage-1bb41.appspot.com"

	return c
}
