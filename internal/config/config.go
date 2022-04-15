package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	AppAddress  string
	KafkaConfig KafkaConfig
}

var config *Config

func Init() {
	mapEnv, err := godotenv.Read("conf/.env")
	if err != nil {
		log.Panicf("[Init] error on reading env: %+v", err)
	}

	config = &Config{
		AppName:    mapEnv["APP_NAME"],
		AppAddress: mapEnv["APP_ADDRESS"],
		KafkaConfig: KafkaConfig{
			Address: mapEnv["KAFKA_ADDRESS"],
		},
	}

	if config.AppName == "" {
		log.Panicf("[Init] app name cannot be empty")
	}

	if config.AppAddress == "" {
		log.Panicf("[Init] app address cannot be empty")
	}

	if config.KafkaConfig.Address == "" {
		log.Panicf("[Init] kafka config cannot be empty")
	}
}

func Get() *Config {
	return config
}
