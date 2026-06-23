package config

// TODO(phase2-crud-config): Load configuration from environment.
//
// Task:
//   1. Config struct: Port (default "8080"), DatabaseURL (required).
//   2. Return error if DATABASE_URL is empty.
//
// Learn: fail fast on missing required config.

type Config struct {
	Port        string
	DatabaseURL string
}

func Load() (*Config, error) {
	panic("TODO: implement config.Load")
}
