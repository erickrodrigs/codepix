package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Transaction ...
type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func (transaction *Transaction) isValid() error {
	v := validator.New()
	err := v.Struct(transaction)

	if err != nil {
		fmt.Printf("Error during transaction validation: %s", err.Error())
		return err
	}

	return nil
}

// ParseJSON ...
func (transaction *Transaction) ParseJSON(data []byte) error {
	err := json.Unmarshal(data, transaction)

	if err != nil {
		return err
	}

	err = transaction.isValid()

	if err != nil {
		return err
	}

	return nil
}

// ToJSON ...
func (transaction *Transaction) ToJSON() ([]byte, error) {
	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(transaction)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
