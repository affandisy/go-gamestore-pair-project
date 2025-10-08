package usecase

import (
	"context"
	"gamestore/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, u *domain.Customer) (int64, error)
	GetById(ctx context.Context, id int64) (*domain.Customer, error)
	Update(ctx context.Context, u *domain.Customer) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]domain.Customer, error)
}
