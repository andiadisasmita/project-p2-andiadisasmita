package models

import "time"

// RentalHistory represents the rental_history table
type RentalHistory struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `gorm:"not null" json:"user_id"`
	StockID    uint       `gorm:"not null" json:"stock_id"`
	RentalDate time.Time  `gorm:"autoCreateTime" json:"rental_date"`
	ReturnDate *time.Time `json:"return_date"` // Nullable
	RentalCost float64    `gorm:"not null" json:"rental_cost"`
	Status     string     `gorm:"not null;default:'reserved'" json:"status"` // e.g., reserved, with_user, returned
}
