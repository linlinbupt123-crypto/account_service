package model

import (
	"time"
)

type Transaction struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	TxID         string    `gorm:"unique;size:64" json:"tx_id"`
	UserID       int64     `json:"user_id"`
	Currency     string    `json:"currency"`
	Change       int64     `json:"change"`
	FrozenChange int64     `json:"frozen_change"`
	BizType      string    `gorm:"type:enum('deposit','withdraw','freeze','unfreeze','trade','transfer')" json:"biz_type"`
	BizID        string    `json:"biz_id"`
	BalanceAfter int64     `json:"balance_after"`
	FrozenAfter  int64     `json:"frozen_after"`
	CreatedAt    time.Time `json:"created_at"`
}
