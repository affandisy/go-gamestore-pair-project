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
	Title        string
	TotalSold    int
	TotalRevenue float64
}
