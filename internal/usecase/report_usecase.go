package usecase

import (
	"gamestore/internal/domain"
	"gamestore/internal/repository"
)

type ReportUsecase struct {
	repo *repository.ReportRepository
}

func NewReportUsecase(repo *repository.ReportRepository) *ReportUsecase {
	return &ReportUsecase{repo: repo}
}

func (u *ReportUsecase) GetPurchaseHistory() ([]domain.PurchaseHistory, error) {
	return u.repo.GetCustomerPurchaseHistory()
}

func (u *ReportUsecase) GetBestSellingGames() ([]domain.BestSeller, error) {
	return u.repo.GetBestSellingGames()
}

func (u *ReportUsecase) GetTotalRevenue() (float64, error) {
	return u.repo.GetTotalRevenue()
}
