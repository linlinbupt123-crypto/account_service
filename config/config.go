package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}
type ServerConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	MySQL  MySQLConfig  `yaml:"mysql"`
	Server ServerConfig `yaml:"server"`
}

func InitConfig() error {
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}
	log.Printf("Config loaded: %+v\n", AppConfig)
	return nil
}
