package src

import (
	"context"
	"errors"

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
		if db.orm.Where(entity.User{Email: u.Email}).Take(&entity.User{}).Error == nil {
			return nil, errors.New("user alredy exists")
		}
	}
	return &u, nil
}

func (db PaSQLRepository) Update(ctx context.Context, u entity.User) (*entity.User, error) {
	err := db.orm.Model(&u).Where("id = ?", u.ID).Updates(entity.User{FirstName: u.FirstName, LastName: u.LastName, Email: u.Email}).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db PaSQLRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	var u entity.User
	err := db.orm.Where("id = ?", id).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db PaSQLRepository) GetAll(ctx context.Context) (*[]entity.User, error) {
	var u []entity.User
	err := db.orm.Find(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db PaSQLRepository) Delete(ctx context.Context, email string) error {
	err := db.orm.Where("email = ?", email).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
