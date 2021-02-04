package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
	"time"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json:"code" valid:""notnull`
	Name string `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil;
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	err := bank.isValid()

	if err != nil {
		return nil, error
	}

	return &bank, nil
}