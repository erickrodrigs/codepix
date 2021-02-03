package usecase

import (
	"errors"
	"log"

	"github.com/erickrodrigs/codepix/codepix-go/domain/model"
)

// TransactionUseCase ...
type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyRepositoryInterface
}

// Register ...
func (useCase *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := useCase.PixRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := useCase.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	useCase.TransactionRepository.Save(transaction)

	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

// Confirm ...
func (useCase *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {
	transaction, err := useCase.TransactionRepository.Find(transactionID)

	if err != nil {
		log.Println("Transaction not found", transactionID)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = useCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Complete ...
func (useCase *TransactionUseCase) Complete(transactionID string) (*model.Transaction, error) {
	transaction, err := useCase.TransactionRepository.Find(transactionID)

	if err != nil {
		log.Println("Transaction not found", transactionID)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = useCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Error ...
func (useCase *TransactionUseCase) Error(transactionID string, reason string) (*model.Transaction, error) {
	transaction, err := useCase.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = useCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
