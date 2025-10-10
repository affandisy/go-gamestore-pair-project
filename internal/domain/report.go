package domain

import "time"

type PurchaseHistory struct {
	OrderID       int64
	CustomerName  string
	CustomerEmail string
	GameTitle     string
	Price         float64
	PaymentStatus string
	OrderDate     time.Time
}

type BestSeller struct {
	GameName        string
	TotalTerjual    int
	TotalPendapatan float64
}

type RevenueSummary struct {
	TotalRevenue     float64
	OutstandingBills float64
	DailyIncome      float64
}

type AdminSummary struct {
	TotalCustomers int
	TotalGames     int
	TotalOrders    int
	TotalPayments  int
	TotalRevenue   float64
}
