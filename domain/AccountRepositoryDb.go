package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/sanvidhans/bankingapp/errs"
	"github.com/sanvidhans/bankingapp/logger"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d *AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES(?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account "+err.Error())
		return nil, errs.NewUnExpectedError("Unexpected error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting lastInsertedId "+err.Error())
		return nil, errs.NewUnExpectedError("Unexpected error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil

}

func NewAccountRepositoryDb (dbclinet *sqlx.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{dbclinet}
}