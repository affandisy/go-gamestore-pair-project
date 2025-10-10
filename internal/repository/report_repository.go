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
	query := `SELECT nama_game, total_terjual, total_pendapatan
		FROM v_best_selling_games
		ORDER BY total_terjual DESC;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bests []domain.BestSeller
	for rows.Next() {
		var b domain.BestSeller
		err := rows.Scan(&b.GameName, &b.TotalTerjual, &b.TotalPendapatan)
		if err != nil {
			return nil, err
		}
		bests = append(bests, b)
	}
	return bests, nil
}

// Total Revenue
func (r *ReportRepository) GetRevenueSummary() (domain.RevenueSummary, error) {
	query := `SELECT total_revenue, outstanding_bills, daily_income
		FROM v_total_revenue;`
	var s domain.RevenueSummary
	err := r.DB.QueryRow(query).Scan(&s.TotalRevenue, &s.OutstandingBills, &s.DailyIncome)
	return s, err
}

// Summary
func (r *ReportRepository) GetAdminSummary() (domain.AdminSummary, error) {
	query := `SELECT total_customers, total_games, total_orders, total_payments, total_revenue
		FROM v_summary;`

	var summary domain.AdminSummary
	err := r.DB.QueryRow(query).Scan(
		&summary.TotalCustomers,
		&summary.TotalGames,
		&summary.TotalOrders,
		&summary.TotalPayments,
		&summary.TotalRevenue,
	)
	return summary, err
}
