package domain

import (
	"github.com/sanvidhans/bankingapp/dto"
	"github.com/sanvidhans/bankingapp/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c *Customer) SetStatusAsText() string  {
	setStatusText := "active"

	if c.Status == "0" {
		setStatusText = "inactive"
	}

	return setStatusText
}

func (c *Customer) ToDto() dto.CustomerResponse {

	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.SetStatusAsText(),
	}

	return response
}