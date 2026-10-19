package config

import "os"

type Config struct {
	Addr         string
	DatabasePath string
	JWTSecret    string
}

func Load() Config {
	return Config{
		Addr:         env("ADDR", ":3000"),
		DatabasePath: env("DATABASE_PATH", "data/conduit.db"),
		JWTSecret:    env("JWT_SECRET", "dev-secret-pls-change-me"),
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
