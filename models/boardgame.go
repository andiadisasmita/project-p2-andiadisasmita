package models

import "time"

// Boardgame represents the boardgames table
type Boardgame struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Availability uint      `gorm:"not null" json:"availability"`
	RentalCost   float64   `gorm:"not null" json:"rental_cost"`
	CategoryID   uint      `gorm:"not null" json:"category_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
