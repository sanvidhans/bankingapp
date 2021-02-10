package service

import (
	"github.com/sanvidhans/bankingapp/domain"
	"github.com/sanvidhans/bankingapp/dto"
	"github.com/sanvidhans/bankingapp/errs"
)

type CustomerService interface {
	GetAllCustomer(string)([]domain.Customer, *errs.AppError)
	GetCustomer(string)(*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func(d DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError){
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	}else{
		status = ""
	}
	return d.repo.FindAll(status)
}

func(d DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError){
	c, err := d.repo.ById(id)
	if err != nil{
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}


func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}