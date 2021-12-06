package config

import (
	"golang.org/x/oauth2/jwt"
)

type Config struct {
	Server ServerConfig `yaml:"database"`
	Conf   *jwt.Config
}

type ServerConfig struct {
	DatabaseUrl string `yaml:"database_url"`
	BucketName  string
}

//type DatabaseConfig struct {
//	DBName     string
//	DBUser     string
//	DBPassword string
//}
//
//const (
//	dbCredentialsFormat = "user=%s password=%s dbname=%s host=%s port=%d sslmode=%s"
//	configFileName      = "config-backendEkskul-%s-.yml"
//)
