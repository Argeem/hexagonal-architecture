package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	as service.AccountService
}

func NewAccountHandler(as service.AccountService) accountHandler {
	return accountHandler{as: as}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	if r.Header.Get("content-type") != "application/json" {
		handlerError(w, errs.NewValidationError("request body incorrect format"))
		return
	}
	req := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handlerError(w, errs.NewValidationError("request body incorrect format"))
		return
	}
	res, err := h.as.NewAccount(customerID, req)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func (h accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.as.GetAccount(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
