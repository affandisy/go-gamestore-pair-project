package usecase

import (
	"gamestore/internal/domain"
	"time"
)

// Kontrak CRUD Games
type GameRepository interface {
	Create(game *domain.Game) error
	FindAll() ([]domain.Game, error)
	FindById(id int64) (*domain.Game, error)
	FindByCategoryID(categoryID int64) ([]domain.Game, error)
	Update(game *domain.Game) error
	Delete(id int64) error
}

type GameUsecase struct {
	repo GameRepository
}

func NewGameUsecase(repo GameRepository) *GameUsecase {
	return &GameUsecase{repo: repo}
}

func (u *GameUsecase) CreateGame(categoryID int64, titles string, price float64) error {
	game := &domain.Game{
		CategoryID: categoryID,
		Title:      titles,
		Price:      price,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	return u.repo.Create(game)
}

func (u *GameUsecase) FindAllGame() ([]domain.Game, error) {
	return u.repo.FindAll()
}

func (u *GameUsecase) FindGameById(id int64) (*domain.Game, error) {
	return u.repo.FindById(id)
}

func (u *GameUsecase) UpdateGame(game *domain.Game) error {
	game.Updated_at = time.Now()
	return u.repo.Update(game)
}

func (u *GameUsecase) DeleteGame(id int64) error {
	return u.repo.Delete(id)
}

func (u *GameUsecase) FindGameByCategoryID(categoryID int64) ([]domain.Game, error) {
	return u.repo.FindByCategoryID(categoryID)
}
