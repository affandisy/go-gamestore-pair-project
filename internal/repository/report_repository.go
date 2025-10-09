package repository

import (
	"database/sql"
	"gamestore/internal/domain"
)

type ReportRepository struct {
	DB *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

// Customer Purchase History
func (r *ReportRepository) GetCustomerPurchaseHistory() ([]domain.PurchaseHistory, error) {
	query := `SELECT OrderID, customer_name, customer_email, game_title, price, payment_status, order_date
		FROM v_customer_purchase_history
		ORDER BY OrderID;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var histories []domain.PurchaseHistory
	for rows.Next() {
		var h domain.PurchaseHistory
		err := rows.Scan(&h.OrderID, &h.CustomerName, &h.CustomerEmail, &h.GameTitle, &h.Price, &h.PaymentStatus, &h.OrderDate)
		if err != nil {
			return nil, err
		}
		histories = append(histories, h)
	}
	return histories, nil
}

// Game Terlaris
func (r *ReportRepository) GetBestSellingGames() ([]domain.BestSeller, error) {
	query := `SELECT title, total_sold, total_revenue
		FROM v_best_selling_games
		ORDER BY total_sold DESC;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bests []domain.BestSeller
	for rows.Next() {
		var b domain.BestSeller
		err := rows.Scan(&b.Title, &b.TotalSold, &b.TotalRevenue)
		if err != nil {
			return nil, err
		}
		bests = append(bests, b)
	}
	return bests, nil
}

func (r *ReportRepository) GetTotalRevenue() (float64, error) {
	query := `SELECT COALESCE(total_revenue, 0) FROM v_total_revenue;`
	var total float64
	err := r.DB.QueryRow(query).Scan(&total)
	return total, err
}
