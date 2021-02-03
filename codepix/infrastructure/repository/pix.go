package repository

import (
	"fmt"

	"github.com/erickrodrigs/codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

// PixKeyRepositoryDb ...
type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

// AddBank ...
func (repo PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := repo.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

// AddAccount ...
func (repo PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := repo.Db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

// RegisterKey ...
func (repo PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := repo.Db.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

// FindKeyByKind ...
func (repo PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	repo.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

// FindAccount ...
func (repo PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	repo.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}

	return &account, nil
}

// FindBank ...
func (repo PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	repo.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}

	return &bank, nil
}
