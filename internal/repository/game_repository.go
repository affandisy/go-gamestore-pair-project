package repository

import (
	"database/sql"
	"errors"
	"gamestore/internal/domain"
)

type GameRepository struct {
	DB *sql.DB
}

func (r *GameRepository) Create(game *domain.Game) error {
	query := `INSERT INTO games (categoryid, titles, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING gameid;`

	err := r.DB.QueryRow(
		query,
		game.CategoryID,
		game.Title,
		game.Price,
		game.Created_at,
		game.Updated_at,
	).Scan(&game.GameID)

	if err != nil {
		return err
	}

	return nil
}

func (r *GameRepository) FindAll() ([]domain.Game, error) {
	query := `SELECT gameid, categoryid, titles, price, created_at, updated_at
		FROM games;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var games []domain.Game
	for rows.Next() {
		var g domain.Game
		err := rows.Scan(
			&g.GameID,
			&g.CategoryID,
			&g.Title,
			&g.Price,
			&g.Created_at,
			&g.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *GameRepository) FindById(id int64) (*domain.Game, error) {
	query := `SELECT gameid, categoryid, titles, price, created_at, updated_at
		FROM games
		WHERE gameid = $1;`

	var g domain.Game
	err := r.DB.QueryRow(query, id).Scan(
		&g.GameID,
		&g.CategoryID,
		&g.Title,
		&g.Price,
		&g.Created_at,
		&g.Updated_at,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &g, nil
}

func (r *GameRepository) Update(game *domain.Game) error {
	query := `UPDATE games
		SET category_id = $1, title = $2, price = $3, updated_at = $4
		WHERE id = $5;`

	result, err := r.DB.Exec(
		query,
		game.CategoryID,
		game.Title,
		game.Price,
		game.Updated_at,
		game.GameID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}

func (r *GameRepository) Delete(id int64) error {
	query := `DELETE FROM games WHERE id = $1;`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}

func (r *GameRepository) FindByCategoryID(categoryID int64) ([]domain.Game, error) {
	query := `SELECT *
	          FROM games
	          WHERE categoryid = $1;`

	rows, err := r.DB.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []domain.Game
	for rows.Next() {
		var g domain.Game
		err := rows.Scan(
			&g.GameID,
			&g.CategoryID,
			&g.Title,
			&g.Price,
			&g.Created_at,
			&g.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		games = append(games, g)
	}

	return games, nil
}
