package domain

import (
	"Banking/dto"
	"Banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (customer Customer) statusAsText() string {
	statusAsText := "active"
	if customer.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (customer Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
