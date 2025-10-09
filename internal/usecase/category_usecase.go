package usecase

import (
	"errors"
	"gamestore/internal/domain"
	"strings"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	FindAll() ([]domain.Category, error)
	FindById(id int64) (*domain.Category, error)
	Update(Category *domain.Category) error
	Delete(id int64) error
}

type CategoryUsecase struct {
	repo CategoryRepository
}

func NewCategoryUsecase(repo CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) CreateCategory(name string) error {
	name = strings.TrimSpace(name)
	category := &domain.Category{
		Name: name,
	}

	return u.repo.Create(category)
}

func (u *CategoryUsecase) FindAllCategories() ([]domain.Category, error) {
	return u.repo.FindAll()
}

func (u *CategoryUsecase) FindCatgeoryById(id int64) (*domain.Category, error) {
	return u.repo.FindById(id)
}

func (u *CategoryUsecase) UpdateCategory(category *domain.Category) error {
	if category == nil {
		return errors.New("category cannot be nil")
	}
	return u.repo.Update(category)
}

func (u *CategoryUsecase) DeleteCategory(id int64) error {
	return u.repo.Delete(id)
}
