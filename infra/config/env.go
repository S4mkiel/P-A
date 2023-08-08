package config

import (
	env "github.com/Netflix/go-env"
	cfg "github.com/S4mkiel/p-a/infra/db/config"
	"github.com/S4mkiel/p-a/infra/http/fiber"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("config",
	fx.Provide(NewConfig),
	fx.Provide(func(cfg Config) fiber.Config { return cfg.Http }),
	fx.Provide(func(cfg Config) cfg.Config { return cfg.Db }),
)

type Config struct {
	Extras *env.EnvSet
	Http   fiber.Config
	Db     cfg.Config
}

func NewConfig(logger *zap.SugaredLogger) Config {
	var cfg Config = Config{}
	err := cfg.loadConfig()
	if err != nil {
		logger.Error(err)
	}

	return cfg
}

func (c *Config) loadConfig() error {
	es, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		return err
	}
	c.Extras = &es

	return nil
}
