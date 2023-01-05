package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DatabaseConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type ParseConfig interface {
	MySqlConfig() *DatabaseConfig
	PostgresConfig() *DatabaseConfig
}

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) MySqlConfig() *DatabaseConfig {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}
	return &DatabaseConfig{
		Driver:   os.Getenv("MS_DRIVER"),
		Username: os.Getenv("MS_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MS_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
}

func (c *Config) PostgresConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	return &DatabaseConfig{
		Driver:   os.Getenv("PS_DRIVER"),
		Username: os.Getenv("PS_USER"),
		Password: os.Getenv("PS_PASSWORD"),
		Port:     os.Getenv("PS_PORT"),
		Database: os.Getenv("PS_NAME"),
	}
}

//func FetchConfig() *DatabaseConfig{
//	driver := flag.String("database", "mysql", "Database to use in backend")
//	flag.Parse()
//	if *driver == "mysql" {
//		return MySqlConfig()
//	}
//	return PostgresConfig()
//}
