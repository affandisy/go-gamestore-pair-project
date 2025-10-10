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
	query := `INSERT INTO payments (CustomerID, Amount, Status, CreatedAt)
		VALUES ($1, $2, $3, $4)
		RETURNING PaymentID;`

	err := r.DB.QueryRow(
		query,
		payment.CustomerID,
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
	query := `SELECT PaymentID, CustomerID, Amount, Status, CreatedAt
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
			&p.CustomerID,
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
	query := `SELECT PaymentID, CustomerID, Amount, Status, CreatedAt
		FROM payments WHERE PaymentID = $1;`

	var p domain.Payment
	err := r.DB.QueryRow(query, id).Scan(
		&p.PaymentID,
		&p.CustomerID,
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
		SET CustomerID = $1, Amount = $2, Status = $3
		WHERE PaymentID = $4;`

	result, err := r.DB.Exec(
		query,
		payment.CustomerID,
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

func (r *PaymentRepository) PayAllGames(payment *domain.Payment) error {
	query := `
		SELECT SUM(g.price)
		FROM orders o
		JOIN games g
		ON o.gameid = g.gameid
		JOIN customers c
		ON o.customerid = c.customerid
		WHERE o.status = 'UNPAID' AND c.customerid = $1
		;
	`

	rows, err := r.DB.Query(query, payment.CustomerID)
	if err != nil {
		return err
	}

	for rows.Next() {
		var p domain.Payment
		err := rows.Scan(&p.Amount)
		if err != nil {
			return err
		}
	}

	query2 := `INSERT INTO payments (CustomerID, Amount, Status, CreatedAt)
		VALUES ($1, $2, $3, $4)
		RETURNING PaymentID;`

	err2 := r.DB.QueryRow(
		query2,
		payment.CustomerID,
		payment.Amount,
		payment.Status,
		payment.CreatedAt,
	).Scan(&payment.PaymentID)

	if err2 != nil {
		return err2
	}

	return nil

}
