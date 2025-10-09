package repository

import (
	"database/sql"
	"errors"
	"gamestore/internal/domain"
)

type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) Create(order *domain.Order) (*domain.Order, error) {
	query := `INSERT INTO orders (CustomerID, GameID, CreatedAt)
		VALUES ($1, $2, $3)
		RETURNING OrderID;`

	err := r.DB.QueryRow(
		query,
		order.CustomerID,
		order.GameID,
		order.CreatedAt,
	).Scan(&order.OrderID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) FindAll() ([]domain.Order, error) {
	query := `SELECT OrderID, CustomerID, GameID, CreatedAt
		FROM orders;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		var o domain.Order
		err := rows.Scan(
			&o.OrderID,
			&o.CustomerID,
			&o.GameID,
			&o.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) FindById(id int64) (*domain.Order, error) {
	query := `SELECT OrderID, CustomerID, GameID, CreatedAt
		FROM orders WHERE OrderID = $1;`

	var o domain.Order
	err := r.DB.QueryRow(query, id).Scan(
		&o.OrderID,
		&o.CustomerID,
		&o.GameID,
		&o.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &o, nil
}

func (r *OrderRepository) Update(order *domain.Order) error {
	query := `UPDATE orders SET CustomerID = $1, GameID = $2, CreatedAt = $3
		WHERE OrderID = $4;`

	result, err := r.DB.Exec(
		query,
		order.CustomerID,
		order.GameID,
		order.OrderID,
		order.CreatedAt,
		order.OrderID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func (r *OrderRepository) Delete(id int64) error {
	query := `DELETE FROM orders WHERE OrderID = $1;`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func (r *OrderRepository) FindAllByCustomerID(customerID int64) ([]domain.Order, error) {
	query := `SELECT o.orderid, g.title, g.price
	          FROM orders o
			  JOIN games g
			  ON o.gameid = g.gameid
	          WHERE o.customerid = $1;`

	rows, err := r.DB.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var g domain.Order
		err := rows.Scan(
			&g.OrderID,
			&g.GameTitle,
			&g.GamePrice,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, g)
	}

	return orders, nil
}
