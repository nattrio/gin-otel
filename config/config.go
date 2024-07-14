package config

import (
	"fmt"
	"log/slog"
	"sync"

	env "github.com/caarlos0/env/v11"
)

var (
	val  *Config
	once sync.Once
)

type Config struct {
	App struct {
		Port string `env:"APP_PORT" envDefault:"8080"`
	}
	Postgres struct {
		Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
		Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
		Username string `env:"POSTGRES_USERNAME" envDefault:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"password"`
		Database string `env:"POSTGRES_DB" envDefault:"postgres"`
	}
}

func (c *Config) PostgresURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Database,
	)
}

func newConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		slog.Warn(err.Error())
	}
	return &cfg
}

func Value() *Config {
	return val
}

func init() {
	once.Do(func() {
		val = newConfig()
	})
}
