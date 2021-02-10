package dto

import "github.com/sanvidhans/bankingapp/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

func (n *NewAccountRequest) Validate() *errs.AppError {
	if n.Amount < 5000 {
		return errs.NewUnProcessableInputs("To open account you must deposit 5000")
	}

	if n.AccountType == "saving" || n.AccountType == "current" {
		return nil
	}else {
		return errs.NewUnProcessableInputs("The Account Type should be saving or current")
	}

}