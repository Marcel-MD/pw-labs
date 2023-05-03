package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	BotToken   string `env:"BOT_TOKEN"`
	Port       string `env:"PORT" envDefault:"8443"`
	BotUrl     string `env:"BOT_URL" envDefault:"https://api.telegram.org/bot"`
	NewsUrl    string `env:"NEWS_URL" envDefault:"https://newsapi.org/v2"`
	NewsApiKey string `env:"NEWS_API_KEY"`
}

func New() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg
}
