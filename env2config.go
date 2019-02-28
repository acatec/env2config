package env2config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type ConfigInterface interface {
}

func PrepareConfig(cfg ConfigInterface, filenames ...string) (err error) {
	if err := godotenv.Load(filenames...); err != nil {
		log.Println("File .env not found, reading configuration from ENV")
	}

	if err := env.Parse(cfg); err != nil {
		log.Fatalln("Failed to parse ENV")
	}

	return
}
