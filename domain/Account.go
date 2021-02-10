package domain

import (
	"github.com/sanvidhans/bankingapp/dto"
	"github.com/sanvidhans/bankingapp/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(a Account)(*Account, *errs.AppError)
}

func (a *Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	resp := dto.NewAccountResponse{
		AccountId: a.AccountId,
	}

	return resp
}