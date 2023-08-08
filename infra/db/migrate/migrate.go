package table

import (
	"github.com/S4mkiel/p-a/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Base{},
		&entity.User{},
	)
	return err
}
