package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/S4mkiel/p-a/domain/module"
	"github.com/S4mkiel/p-a/infra/config"
	"github.com/S4mkiel/p-a/infra/db"
	"github.com/S4mkiel/p-a/infra/http/fiber"
	logger "github.com/S4mkiel/p-a/infra/log"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if os.Getenv("ENV") != "production" {
		LoadConfig()
	}
	fx.New(
		fiber.Module,
		config.Module,
		logger.Module,
		db.Module,

		module.Service,
	).Run()
}

func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}
