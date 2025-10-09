package repository

import (
	"database/sql"
	"errors"
	"gamestore/internal/domain"
)

type CustomerRepository struct {
	DB *sql.DB
}

func (r *CustomerRepository) Create(customer *domain.Customer) error {
	query := `INSERT INTO customers (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING customer_id;`

	return r.DB.QueryRow(
		query,
		customer.Name,
		customer.Email,
		customer.Password,
		customer.CreatedAt,
		customer.UpdatedAt,
	).Scan(&customer.CustomerID)
}

func (r *CustomerRepository) FindAll() ([]domain.Customer, error) {
	query := `SELECT customer_id, name, email, password, created_at, updated_at
		FROM customers;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []domain.Customer

	for rows.Next() {
		var c domain.Customer
		err := rows.Scan(
			&c.CustomerID,
			&c.Name,
			&c.Email,
			&c.Password,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}

func (r *CustomerRepository) FindById(id int64) (*domain.Customer, error) {
	query := `SELECT customer_id, name, email, password, created_at, updated_at
		FROM customers WHERE customer_id = $1;`

	var c domain.Customer

	err := r.DB.QueryRow(query, id).Scan(
		&c.CustomerID,
		&c.Name,
		&c.Email,
		&c.Password,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *CustomerRepository) Update(customer *domain.Customer) error {
	query := `UPDATE customers
		SET name = $1, email = $2, password = $3, updated_at = $4
		WHERE customer_id = $5;`

	res, err := r.DB.Exec(
		query,
		customer.Name,
		customer.Email,
		customer.Password,
		customer.UpdatedAt,
		customer.CustomerID,
	)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("customer not found")
	}
	return nil
}

func (r *CustomerRepository) Delete(id int64) error {
	query := `DELETE FROM customers WHERE customer_id = $1;`

	res, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("customer not found")
	}

	return nil
}
