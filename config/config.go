package config

import (
	"flyinghero/nes"
	"os"
	"time"
)

type Config struct {
	URL      string
	Interval time.Duration
	NESKeys  *nes.NESKeys
}

func Load() *Config {
	nesKeys := nes.GetNESKeys()

	return &Config{
		URL:      getEnv("GAME_URL", "https://example.com"),
		Interval: getDurationFromEnv("INTERVAL", 5*time.Second),
		NESKeys:  nesKeys,
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
