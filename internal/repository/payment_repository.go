package repository

import (
	"database/sql"
	"errors"
	"gamestore/internal/domain"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) Create(payment *domain.Payment) error {
	query := `INSERT INTO payments (OrderID, Amount, Status, CreatedAt)
		VALUES ($1, $2, $3, $4)
		RETURNING PaymentID;`

	err := r.DB.QueryRow(
		query,
		payment.OrderID,
		payment.Amount,
		payment.Status,
		payment.CreatedAt,
	).Scan(&payment.PaymentID)

	if err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepository) FindAll() ([]domain.Payment, error) {
	query := `SELECT PaymentID, OrderID, Amount, Status, CreatedAt
		FROM payments;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var payments []domain.Payment

	for rows.Next() {
		var p domain.Payment
		err := rows.Scan(
			&p.PaymentID,
			&p.OrderID,
			&p.Amount,
			&p.Status,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *PaymentRepository) FindById(id int64) (*domain.Payment, error) {
	query := `SELECT PaymentID, OrderID, Amount, Status, CreatedAt
		FROM payments WHERE PaymentID = $1;`

	var p domain.Payment
	err := r.DB.QueryRow(query, id).Scan(
		&p.PaymentID,
		&p.OrderID,
		&p.Amount,
		&p.Status,
		&p.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *PaymentRepository) Update(payment *domain.Payment) error {
	query := `UPDATE payments
		SET OrderID = $1, Amount = $2, Status = $3
		WHERE PaymentID = $4;`

	result, err := r.DB.Exec(
		query,
		payment.OrderID,
		payment.Amount,
		payment.Status,
		payment.PaymentID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("payment not found")
	}

	return nil
}

func (r *PaymentRepository) Delete(id int64) error {
	query := `DELETE FROM payments WHERE PaymentID = $1;`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("payment not found")
	}

	return nil
}
