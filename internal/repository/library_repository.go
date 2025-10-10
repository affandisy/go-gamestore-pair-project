package repository

import (
	"database/sql"
	"errors"
	"gamestore/internal/domain"
)

type LibraryRepository struct {
	DB *sql.DB
}

func (r *LibraryRepository) Create(library *domain.Library) error {
	query := `
		INSERT INTO library (customerid, gameid) VALUES
		($1, $2) RETURNING libraryid;
	`

	err := r.DB.QueryRow(query, library.CustomerID, library.GameID).Scan(&library.LibraryID)
	if err != nil {
		return err
	}

	return nil
}

func (r *LibraryRepository) FindAllUserGame(customerID int64) ([]domain.Library, error) {
	query := `
		SELECT l.libraryid, l.gameid, g.title, l.createdat
		FROM library l
		JOIN games g
		ON l.gameid = g.gameid
		WHERE customerid = $1
		;
	`

	rows, err := r.DB.Query(query, customerID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var library []domain.Library

	for rows.Next() {
		var l domain.Library
		err := rows.Scan(
			&l.LibraryID,
			&l.GameID,
			&l.GameTitle,
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		library = append(library, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return library, nil
}

func (r *LibraryRepository) FindById(id int64) (*domain.Library, error) {
	query := `SELECT libraryid, CustomerID, GameID, CreatedAt
		FROM library WHERE libraryid = $1;`

	var l domain.Library
	err := r.DB.QueryRow(query, id).Scan(
		&l.LibraryID,
		&l.CustomerID,
		&l.GameID,
		&l.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &l, nil
}

func (r *LibraryRepository) FindByGameId(customerID, gameID int64) (*domain.Library, error) {
	query := `SELECT libraryid, CustomerID, GameID, CreatedAt
		FROM library WHERE customerid = $1 AND gameid = $2;`

	var l domain.Library
	err := r.DB.QueryRow(query, customerID, gameID).Scan(
		&l.LibraryID,
		&l.CustomerID,
		&l.GameID,
		&l.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &l, nil
}
