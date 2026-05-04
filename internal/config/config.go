package config

import "os"

type Config struct{
	DatabaseURL string;
	RedisURL string;
	Port string;
	BaseURL string;
}

func Load() *Config{
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL: os.Getenv("REDIS_URL"),
		Port: os.Getenv("Port"),
		BaseURL: os.Getenv("BASE_URL"),
	}
}
