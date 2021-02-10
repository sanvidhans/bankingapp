package service

import (
	"github.com/sanvidhans/bankingapp/domain"
	"github.com/sanvidhans/bankingapp/dto"
	"github.com/sanvidhans/bankingapp/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)  {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	newAccount := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: "saving",
		Amount:      req.Amount,
		Status:      "1",
	}
	result, err := d.repo.Save(newAccount)
	if err != nil {
		return nil, err
	}

	response := result.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}