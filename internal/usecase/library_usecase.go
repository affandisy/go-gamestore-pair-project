package usecase

import (
	"gamestore/internal/domain"
)

type LibraryRepository interface {
	Create(library *domain.Library) error
	FindAllUserGame(customerID int64) ([]domain.Library, error)
	FindById(id int64) (*domain.Library, error)
	// Delete(id int64) error
}

type Libraryusecase struct {
	repo LibraryRepository
}

func NewLibraryUsecase(repo LibraryRepository) *Libraryusecase {
	return &Libraryusecase{repo: repo}
}

func (u *Libraryusecase) CreateGameInLibrary(customerID, gameID int64) error {
	library := &domain.Library{
		CustomerID: customerID,
		GameID:     gameID,
	}

	return u.repo.Create(library)
}

func (u *Libraryusecase) FindAllGamesInLibrary(customerID int64) ([]domain.Library, error) {
	return u.repo.FindAllUserGame(customerID)
}

func (u *Libraryusecase) FindGameInLibraryByID(id int64) (*domain.Library, error) {
	return u.repo.FindById(id)
}

// func (u *Libraryusecase) DeleteGameInLibrary(id int64) error {
// 	return u.repo.Delete(id)
// }
