package usecase

import (
	"errors"

	"github.com/erickrodrigs/codepix/codepix-go/domain/model"
)

// PixUseCase ...
type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

// RegisterKey ...
func (useCase *PixUseCase) RegisterKey(key string, kind, string, accountID string) (*model.PixKey, error) {
	account, err := useCase.PixKeyRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)

	if err != nil {
		return nil, err
	}

	useCase.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

// FindKey ...
func (useCase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := useCase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
