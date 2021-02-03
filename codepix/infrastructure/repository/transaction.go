package repository

import (
	"fmt"

	"github.com/erickrodrigs/codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

// TransactionRepositoryDb ...
type TransactionRepositoryDb struct {
	Db *gorm.DB
}

// Register ...
func (repo *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := repo.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

// Save ...
func (repo *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := repo.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

// Find ...
func (repo *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	repo.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
