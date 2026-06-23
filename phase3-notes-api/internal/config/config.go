package config

// TODO(phase3-config): Load all configuration for API and worker.
//
// Task:
//   1. Config struct fields:
//        Port, DatabaseURL, RedisURL, SessionSecret, WorkerPollInterval
//   2. Load from env with defaults where sensible.
//   3. Validate required fields (DatabaseURL, SessionSecret).
//   4. Share this package between cmd/api and cmd/worker.
//
// Learn: one config package, loaded at composition root only.

type Config struct {
	Port               string
	DatabaseURL        string
	RedisURL           string
	SessionSecret      string
	WorkerPollInterval string
}

func Load() (*Config, error) {
	panic("TODO: implement config.Load")
}
