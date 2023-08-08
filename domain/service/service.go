package service

import (
	"context"

	"github.com/S4mkiel/p-a/domain/entity"
	"github.com/S4mkiel/p-a/domain/repository"
)

type PaService struct {
	repo repository.PaRepository
}

func NewPaService(repo repository.PaRepository) *PaService {
	return &PaService{
		repo: repo,
	}
}

func (s *PaService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	u, e := s.repo.Create(ctx, user)
	if e != nil {
		return nil, e
	} else {
		return u, nil
	}
}


func (s *PaService) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	u, e := s.repo.Update(ctx, user)
	if e != nil {
		return nil, e
	} else {
		return u, nil
	}
}

func (s *PaService) Get(ctx context.Context, id int) (*entity.User, error) {
	u, e := s.repo.Get(ctx, id)
	if e != nil {
		return nil, e
	} else {
		return u, nil
	}
}

func (s *PaService) GetAll(ctx context.Context) (*[]entity.User, error) {
	u, e := s.repo.GetAll(ctx)
	if e != nil {
		return nil, e
	} else {
		return u, nil
	}
}

func (s *PaService) Delete(ctx context.Context, user string) error {
	e := s.repo.Delete(ctx, user)
	if e != nil {
		return e
	} else {
		return nil
	}
}