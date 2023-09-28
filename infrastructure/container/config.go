package container

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/palantir/stacktrace"
	"time"
)

type Config struct {
	DatabaseDSN       string        `required:"true" envconfig:"DATABASE_DSN"`
	HTTPTimeout       int           `default:"30" envconfig:"HTTP_TIMEOUT"`
	HTTPClientTimeout time.Duration `default:"30s" envconfig:"HTTP_CLIENT_TIMEOUT"`
	HTTPServerAddr    string        `default:":8080" envconfig:"HTTP_SERVER_ADDR"`

	BCryptCost int `default:"14" envconfig:"BCRYPT_COST"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, stacktrace.NewError("Error loading .env file")
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, stacktrace.Propagate(err, "Error loading config from environment variables")
	}

	return &cfg, nil
}
