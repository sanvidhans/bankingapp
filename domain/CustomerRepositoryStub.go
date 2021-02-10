package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func(c CustomerRepositoryStub) FindAll() ([]Customer, error){
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1234", "Sanvidhan", "Pune", "411014", "19051991", "Active"},
		{"5678", "Sayali", "Pune", "411014", "19051991", "Active"},
		{"91011", "Rahul", "Pune", "411014", "19051991", "Active"},
	}

	return  CustomerRepositoryStub{
		customers: customers,
	}
}
