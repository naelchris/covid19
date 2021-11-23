package config

type Config struct {
	Server   ServerConfig `yaml:"database"`
}

type ServerConfig struct {
	DatabaseUrl string `yaml:"database_url"`
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
