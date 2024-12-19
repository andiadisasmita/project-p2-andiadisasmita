package models

import "time"

// Review represents the reviews table
type Review struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	BoardgameID uint      `gorm:"not null" json:"boardgame_id"`
	Rating      uint      `gorm:"not null" json:"rating"` // 1 to 5
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
