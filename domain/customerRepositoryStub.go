package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Serhii", "Kyiv", "01001", "2008-01-01", "1"},
		{"1002", "Eugene", "Kyiv", "01001", "2008-01-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}
