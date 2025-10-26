package config

import (
	"flyinghero/nes"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	URL         string
	Interval    time.Duration
	NESKeys     *nes.NESKeys
	GameElement string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Файл .env не найден, используем переменные окружения: %v", err)
	}

	nesKeys := nes.GetNESKeys()

	return &Config{
		URL:         getEnv("GAME_URL", "https://example.com"),
		Interval:    getDurationFromEnv("INTERVAL", 5*time.Second),
		NESKeys:     nesKeys,
		GameElement: getEnv("GAME_ELEMENT", "#nesroot"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationFromEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
