package repository

import (
	"database/sql"
	"errors"
	"gamestore/internal/domain"
)

type CategoryRepository struct {
	DB *sql.DB
}

func (r *CategoryRepository) Create(category *domain.Category) error {
	query := `INSERT INTO Categories (Name)
		VALUES ($1)
		RETURNING CategoryID;`

	return r.DB.QueryRow(
		query,
		category.Name,
	).Scan(&category.CategoryID)
}

func (r *CategoryRepository) FindAll() ([]domain.Category, error) {
	query := `SELECT CategoryID, Name
		FROM Categories;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var cat domain.Category
		err := rows.Scan(
			&cat.CategoryID,
			&cat.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, rows.Err()
}

func (r *CategoryRepository) FindById(id int64) (*domain.Category, error) {
	query := `SELECT CategoryID, Name WHERE CategoryID = $1`

	var cat domain.Category

	err := r.DB.QueryRow(query, id).Scan(
		&cat.CategoryID,
		&cat.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cat, nil
}

func (r *CategoryRepository) Update(category *domain.Category) error {
	query := `UPDATE Categories
		SET name = $1 WHERE CategoryID = $2;`

	res, err := r.DB.Exec(
		query,
		category.Name,
		category.CategoryID,
	)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *CategoryRepository) Delete(id int64) error {
	query := `DELETE FROM Categories WHERE CategoryID = $1;`

	res, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("category not found")
	}

	return nil
}
