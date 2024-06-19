package service

import (
	"Banking/domain"
	"Banking/dto"
	"Banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	//dtoCustomers := make([]dto.CustomerResponse, 0)
	//for _, customer := range customers {
	//	dtoCustomers = append(dtoCustomers, customer.ToDto())
	//}

	dtoCustomers := s.ToDtoArray(customers)

	return dtoCustomers, nil
}

func (s DefaultCustomerService) ToDtoArray(customers []domain.Customer) []dto.CustomerResponse {
	dtoCustomers := make([]dto.CustomerResponse, 0)
	for _, customer := range customers {
		dtoCustomers = append(dtoCustomers, customer.ToDto())
	}

	return dtoCustomers
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	var response = customer.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
