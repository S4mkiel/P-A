package src

import (
	"context"

	"github.com/S4mkiel/p-a/domain/entity"
	"gorm.io/gorm"
)

type PaSQLRepository struct {
	orm *gorm.DB
}

func NewPaRepository(db *gorm.DB) *PaSQLRepository {
	return &PaSQLRepository{orm: db}
}

func (db PaSQLRepository) Create(ctx context.Context, u entity.User) (*entity.User, error) {
	err := db.orm.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
