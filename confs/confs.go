package confs

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file")
	}

	return nil
}
