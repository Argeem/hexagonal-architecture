package handler

import "bank/service"

type customerHandler struct {
	cs service.CustomerService
}

func NewCustomerHandler(cs service.CustomerService) customerHandler {
	return customerHandler{cs: cs}
}
