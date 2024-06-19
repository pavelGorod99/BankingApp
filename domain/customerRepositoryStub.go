package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"001", "Pavel Gorodetchi", "Houston", "644617", "1999-06-10", "1"},
		{"002", "John Smith", "New York", "165122", "1990-03-11", "1"},
		{"003", "William Gray", "Washington", "515611", "1991-05-15", "0"},
		{"004", "Andrew Adam", "Los Angeles", "846534", "1995-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
