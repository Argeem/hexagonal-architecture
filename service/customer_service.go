package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	cr repository.CustomerRepository
}

func NewCustomerService(cr repository.CustomerRepository) customerService {
	return customerService{cr: cr}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.cr.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custResS := []CustomerResponse{}
	for _, customer := range customers {
		custRes := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResS = append(custResS, custRes)
	}
	return custResS, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.cr.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custRes := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custRes, nil
}
