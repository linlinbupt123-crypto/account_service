package service

import (
	"errors"

	"github.com/linlinbupt123-crypto/account_service/model"
	"github.com/linlinbupt123-crypto/account_service/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountService struct{}

func (s *AccountService) CreateAccount(userID int64, currency string) error {
	acc := model.Account{
		UserID:    userID,
		Currency:  currency,
		Available: 0,
		Frozen:    0,
	}
	return repository.DB.Create(&acc).Error
}

func (s *AccountService) GetBalance(userID int64, currency string) (*model.Account, error) {
	var acc model.Account
	err := repository.DB.Where("user_id=? AND currency=?", userID, currency).First(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (s *AccountService) FreezeFunds(userID int64, currency string, amount int64, txID, bizID string) error {
	return repository.DB.Transaction(func(tx *gorm.DB) error {
		// 幂等检查
		var exist model.Transaction
		if err := tx.Where("tx_id = ?", txID).First(&exist).Error; err == nil {
			return nil
		}

		var acc model.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id=? AND currency=?", userID, currency).First(&acc).Error; err != nil {
			return err
		}

		if acc.Available < amount {
			return errors.New("insufficient balance")
		}

		acc.Available -= amount
		acc.Frozen += amount

		if err := tx.Save(&acc).Error; err != nil {
			return err
		}

		txRecord := model.Transaction{
			TxID:         txID,
			UserID:       userID,
			Currency:     currency,
			Change:       -amount,
			FrozenChange: amount,
			BizType:      "freeze",
			BizID:        bizID,
			BalanceAfter: acc.Available,
			FrozenAfter:  acc.Frozen,
		}
		return tx.Create(&txRecord).Error
	})
}
