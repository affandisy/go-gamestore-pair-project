package domain

import "time"

type Library struct {
	LibraryID  int64
	CustomerID int64
	GameID     int64
	GameTitle  string
	GamePrice  float64
	CreatedAt  time.Time
}
