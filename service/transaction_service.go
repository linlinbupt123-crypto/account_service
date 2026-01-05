package service

import (
	"errors"

	"github.com/linlinbupt123-crypto/account_service/model"
	"github.com/linlinbupt123-crypto/account_service/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionService struct{}

func (s *TransactionService) Deposit(userID int64, currency string, amount int64, txID, bizID string) error {
	return repository.DB.Transaction(func(tx *gorm.DB) error {
		var exist model.Transaction
		if err := tx.Where("tx_id=?", txID).First(&exist).Error; err == nil {
			return nil
		}

		var acc model.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id=? AND currency=?", userID, currency).First(&acc).Error; err != nil {
			return err
		}

		acc.Available += amount
		if err := tx.Save(&acc).Error; err != nil {
			return err
		}

		txRecord := model.Transaction{
			TxID:         txID,
			UserID:       userID,
			Currency:     currency,
			Change:       amount,
			FrozenChange: 0,
			BizType:      "deposit",
			BizID:        bizID,
			BalanceAfter: acc.Available,
			FrozenAfter:  acc.Frozen,
		}
		return tx.Create(&txRecord).Error
	})
}

func (s *TransactionService) Withdraw(userID int64, currency string, amount int64, txID, bizID string) error {
	return repository.DB.Transaction(func(tx *gorm.DB) error {
		var exist model.Transaction
		if err := tx.Where("tx_id=?", txID).First(&exist).Error; err == nil {
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
		if err := tx.Save(&acc).Error; err != nil {
			return err
		}

		txRecord := model.Transaction{
			TxID:         txID,
			UserID:       userID,
			Currency:     currency,
			Change:       -amount,
			FrozenChange: 0,
			BizType:      "withdraw",
			BizID:        bizID,
			BalanceAfter: acc.Available,
			FrozenAfter:  acc.Frozen,
		}
		return tx.Create(&txRecord).Error
	})
}
