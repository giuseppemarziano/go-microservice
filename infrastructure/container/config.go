package container

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/palantir/stacktrace"
	"time"
)

type Config struct {
	DatabaseUser     string `required:"true" envconfig:"DATABASE_USER"`
	DatabasePassword string `required:"true" envconfig:"DATABASE_PASSWORD"`
	DatabaseHost     string `default:"localhost" envconfig:"DATABASE_HOST"`
	DatabasePort     string `default:"3306" envconfig:"DATABASE_PORT"`
	DatabaseName     string `required:"true" envconfig:"DATABASE_NAME"`
	DatabaseParams   string `default:"?parseTime=true" envconfig:"DATABASE_PARAMS"`

	HTTPTimeout       int           `default:"30" envconfig:"HTTP_TIMEOUT"`
	HTTPClientTimeout time.Duration `default:"30s" envconfig:"HTTP_CLIENT_TIMEOUT"`
	HTTPServerAddr    string        `default:":8080" envconfig:"HTTP_SERVER_ADDR"`

	BCryptCost int `default:"14" envconfig:"BCRYPT_COST"`
}

func (cfg *Config) DatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName, cfg.DatabaseParams)
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
