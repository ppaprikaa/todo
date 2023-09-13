package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DSN string `json:"dsn" yaml:"dsn"`
}

var CFG Config

func init() {
	configPath := os.Getenv("LOS_TODOS_HERMANOS")

	if err := cleanenv.ReadConfig(configPath, &CFG); err != nil {
		log.Fatal(err)
	}
}
