package repository

import (
	"context"

	"github.com/S4mkiel/p-a/domain/entity"
)

type PaRepository interface {
	Create(context.Context, entity.User) (*entity.User, error)
	Update(context.Context, entity.User) (*entity.User, error)
	Get(context.Context, int) (*entity.User, error)
	GetAll(context.Context) (*[]entity.User, error)
	Delete(context.Context, string) error
}
