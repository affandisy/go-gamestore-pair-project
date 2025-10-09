package usecase

import (
	"errors"
	"gamestore/internal/domain"
	"time"
)

type CustomerRepository interface {
	Create(Customer *domain.Customer) error
	FindAll() ([]domain.Customer, error)
	FindById(id int64) (*domain.Customer, error)
	Update(Customer *domain.Customer) error
	Delete(id int64) error
}

type CustomerUsecase struct {
	repo CustomerRepository
}

func NewCustomerUsecase(repo CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

func (u *CustomerUsecase) CreateCustomer(name, email, password string) error {
	customer := &domain.Customer{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.repo.Create(customer)
}

func (u *CustomerUsecase) FindAllCustomer() ([]domain.Customer, error) {
	return u.repo.FindAll()
}

func (u *CustomerUsecase) FindCustomerByID(id int64) (*domain.Customer, error) {
	return u.repo.FindById(id)
}

func (u *CustomerUsecase) UpdateCustomer(customer *domain.Customer) error {
	if customer == nil {
		return errors.New("customer tidak boleh kosong")
	}
	customer.UpdatedAt = time.Now()
	return u.repo.Update(customer)
}

func (u *CustomerUsecase) DeleteCustomer(id int64) error {
	return u.repo.Delete(id)
}

func (u *CustomerUsecase) FindByEmail(email string) (*domain.Customer, error) {
	customers, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, c := range customers {
		if c.Email == email {
			return &c, err
		}
	}

	return nil, nil
}

func (u *CustomerUsecase) Login(email, password string) (*domain.Customer, error) {
	customer, err := u.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("Email belum terdaftar")
	}

	if customer.Password != password {
		return nil, errors.New("Password salah")
	}

	return customer, nil
}
