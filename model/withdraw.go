package model

import (
	"time"
)

type Withdraw struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	UserID   int64  `gorm:"not null;index"`
	Currency string `gorm:"type:varchar(16);not null"`
	Amount   int64  `gorm:"not null"`
	Address  string `gorm:"type:varchar(128);not null"`

	Status string `gorm:"type:varchar(32);not null;index"`
	TxHash string `gorm:"type:varchar(128)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
