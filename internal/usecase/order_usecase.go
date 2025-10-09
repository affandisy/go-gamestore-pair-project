package usecase

import (
	"errors"
	"gamestore/internal/domain"
	"time"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	FindAll() ([]domain.Order, error)
	FindById(id int64) (*domain.Order, error)
	FindAllByCustomerID(customerID int64) ([]domain.Order, error)
	Update(order *domain.Order) error
	Delete(id int64) error
}

type Orderusecase struct {
	repo OrderRepository
}

func NewOrderUsecase(repo OrderRepository) *Orderusecase {
	return &Orderusecase{repo: repo}
}

func (u *Orderusecase) CreateOrder(customerID, gameID int64) error {
	order := &domain.Order{
		CustomerID: customerID,
		GameID:     gameID,
		CreatedAt:  time.Now(),
	}

	return u.repo.Create(order)
}

func (u *Orderusecase) FindAllOrders() ([]domain.Order, error) {
	return u.repo.FindAll()
}

func (u *Orderusecase) FindOrderByID(id int64) (*domain.Order, error) {
	return u.repo.FindById(id)
}

func (u *Orderusecase) UpdateOrder(order *domain.Order) error {
	if order == nil {
		return errors.New("order cannot be nil")
	}
	return u.repo.Update(order)
}

func (u *Orderusecase) DeleteOrder(id int64) error {
	return u.repo.Delete(id)
}

func (u *Orderusecase) FindAllOrderByCustomerID(customerID int64) ([]domain.Order, error) {
	return u.repo.FindAllByCustomerID(customerID)
}
