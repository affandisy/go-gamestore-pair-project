package usecase

import (
	"errors"
	"gamestore/internal/domain"
	"strings"
	"time"
)

type PaymentRepository interface {
	Create(payment *domain.Payment) error
	FindAll() ([]domain.Payment, error)
	FindById(id int64) (*domain.Payment, error)
	Update(payment *domain.Payment) error
	Delete(id int64) error
}

type Paymentusecase struct {
	repo PaymentRepository
}

func NewPaymentUsecase(repo PaymentRepository) *Paymentusecase {
	return &Paymentusecase{repo: repo}
}

func (u *Paymentusecase) CreatePayment(orderID int64, amount float64, status string) error {
	payment := &domain.Payment{
		OrderID:   orderID,
		Amount:    amount,
		Status:    status,
		CreatedAt: time.Now(),
	}

	return u.repo.Create(payment)
}

func (u *Paymentusecase) FindAllPayments() ([]domain.Payment, error) {
	return u.repo.FindAll()
}

func (u *Paymentusecase) FindPaymentByID(id int64) (*domain.Payment, error) {
	return u.repo.FindById(id)
}

func (u *Paymentusecase) UpdatePayment(payment *domain.Payment) error {
	if payment == nil {
		return errors.New("payment cannot be nil")
	}
	payment.Status = strings.ToUpper(strings.TrimSpace(payment.Status))
	return u.repo.Update(payment)
}

func (u *Paymentusecase) DeletePayment(id int64) error {
	return u.repo.Delete(id)
}
