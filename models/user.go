package models

import "time"

// User represents the users table
type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"unique;not null" json:"email"`
	Password   string    `gorm:"not null" json:"-"`
	DepositAmt float64   `gorm:"default:0" json:"deposit_amt"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
