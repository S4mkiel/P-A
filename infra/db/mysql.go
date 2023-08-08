package db

import (
	"context"
	"fmt"

	cfg "github.com/S4mkiel/p-a/infra/db/config"
	table "github.com/S4mkiel/p-a/infra/db/migrate"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlModule = fx.Module("mysql", fx.Provide(NewMySqlClient), fx.Invoke(HookMysqlDatabase))

func NewMySqlClient(c cfg.Config) *gorm.DB {
	dsn := fmt.Sprint(c.Username, ":", c.Password, "@tcp(", c.Host, ")/", c.Database, "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func HookMysqlDatabase(lc fx.Lifecycle, db *gorm.DB, logger *zap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			dbDriver, err := db.DB()
			if err != nil {
				logger.Fatal("Failed to get DB driver", zap.Error(err))
				return err
			}

			err = dbDriver.Ping()
			if err != nil {
				logger.Fatal("failed to ping database", zap.Error(err))
				return err
			}

			err = table.Migrate(db)
			if err != nil {
				logger.Fatal("failed to migrate database", zap.Error(err))
				return err
			}

			logger.Info("MySQL connection has been established successfully!")
			return nil
		},
		OnStop: func(context.Context) error {
			dbDriver, err := db.DB()
			if err != nil {
				logger.Fatal("Failed to get DB driver", zap.Error(err))

			}

			err = dbDriver.Close()
			if err != nil {
				logger.Fatal("failed to close database connection", zap.Error(err))

			}
			return nil
		},
	})
}
