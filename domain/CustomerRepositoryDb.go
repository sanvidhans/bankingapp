package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sanvidhans/bankingapp/errs"
	"github.com/sanvidhans/bankingapp/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)
	findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"

	if status != "" {
		findAllSql += " WHERE status=?"
		err = c.client.Select(&customers, findAllSql, status)
		//rows, err = c.client.Query(findAllSql, status)
	} else{
		err = c.client.Select(&customers, findAllSql)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("No Customers Found")
		}else{
			logger.Error("Error while fetching customers from DB "+err.Error())
			return nil, errs.NewUnExpectedError("Server Exception")
		}
	}

	return customers, nil
}

func (d *CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError)  {
	var c Customer

	customerSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id=?"

	err := d.client.Get(&c, customerSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found!")
		}else{
			logger.Error("Error while scanning row " + err.Error())
			return nil, errs.NewUnExpectedError("Unexpected database error!")
		}

	}

	return &c, nil
}

func NewCustomerRepositoryDb(client *sqlx.DB) *CustomerRepositoryDb {
	return &CustomerRepositoryDb{client}
}