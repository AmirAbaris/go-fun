package config

import (
	"errors"
	"os"
)

// TODO(phase2-auth-config): Load server configuration from environment.
//
// Task:
//   1. Define Config with Port, SessionSecret (or JWTSecret depending on auth strategy).
//   2. SessionSecret must be at least 32 bytes — return error if too short.
//   3. Load from env: PORT, SESSION_SECRET (or JWT_SECRET).
//
// Learn: secrets belong in env vars, never in source code.

type Config struct {
	Port          string
	SessionSecret string
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	sessionSecret := os.Getenv("SESSION_SECRET")

	if port == "" {
		return nil, errors.New("No port bro")
	}

	if sessionSecret == "" {
		return nil, errors.New("No session sec bro")
	}

	if len(sessionSecret) < 32 {
		return nil, errors.New("bad session sec bro")
	}

	return &Config{
		Port:          port,
		SessionSecret: sessionSecret,
	}, nil
}
