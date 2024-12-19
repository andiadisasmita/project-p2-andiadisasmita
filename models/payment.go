package models

import "time"

// Payment represents the payments table
type Payment struct {
	ID       uint       `gorm:"primaryKey" json:"id"`
	RentalID uint       `gorm:"not null" json:"rental_id"`
	Amount   float64    `gorm:"not null" json:"amount"`
	Status   string     `gorm:"not null;default:'unpaid'" json:"status"` // e.g., paid, unpaid
	PaidAt   *time.Time `json:"paid_at"`                                 // Nullable
}
