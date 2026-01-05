package model

import (
	"time"
)

type Account struct {
	UserID    int64     `gorm:"primaryKey" json:"user_id"`
	Currency  string    `gorm:"primaryKey" json:"currency"`
	Available int64     `json:"available"`
	Frozen    int64     `json:"frozen"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
