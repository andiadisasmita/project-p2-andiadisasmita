package models

// Stock represents the stock table
type Stock struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	BoardgameID uint   `gorm:"not null" json:"boardgame_id"`
	Status      string `gorm:"not null;default:'warehouse'" json:"status"` // e.g., warehouse, with_user, to_user, to_warehouse
	Location    string `gorm:"not null" json:"location"`                   // Simplified for warehouse or delivery tracking
}
