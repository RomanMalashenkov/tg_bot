package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type BotConf struct {
	TelegramToken string
	Store         string
}

// получение конфигур бота из переменных окружения
func GetConfig() (BotConf, error) {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env getting from actual env")
	}

	return BotConf{
		TelegramToken: os.Getenv("TELE_TOKEN"),
		Store:         os.Getenv("STORE"),
	}, nil
}
