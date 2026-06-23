package config

// TODO(phase2-todo-config): Load server configuration from environment.
//
// Task:
//   1. Define Config struct with Port (default "8080") and ServerTimeout fields.
//   2. Implement Load() (*Config, error) reading from os.Getenv.
//   3. Use sensible defaults when env vars are unset.
//
// Learn: https://pkg.go.dev/os#Getenv

type Config struct {
	Port string
}

func Load() (*Config, error) {
	panic("TODO: implement config.Load")
}
