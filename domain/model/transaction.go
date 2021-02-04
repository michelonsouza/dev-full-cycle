package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
	"time"
)

const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base `valid:"required"`,
	AccountFrom *Account `valid:"-"`
	Amount float64 `json:"amount" valid:"notnull"` 
	PixKeyTo *PixKey `valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	CancelDescription string `json:"amount" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return error.New("the amount must be greater then 0")
	}

	if t.Status != Transaction.pending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return error.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return error.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	err := transaction.isValid()

	if err != nil {
		return nil, error
	}

	return &transaction, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err;
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.Description = description
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err;
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err;
}