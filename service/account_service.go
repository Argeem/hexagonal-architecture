package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"strings"
	"time"
)

type accountService struct {
	ar repository.AccountRepository
}

func NewAccountService(ar repository.AccountRepository) AccountService {
	return accountService{ar: ar}
}

func (s accountService) NewAccount(customerID int, r NewAccountRequest) (*AccountResponse, error) {
	// validate input
	if r.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checkig")
	}

	// map view to model
	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"), //day of week 0 month 1 day 2 hour 3 minute 4 second 5 year 6
		AccountType: r.AccountType,
		Amount:      r.Amount,
		Status:      1,
	}
	newAcc, err := s.ar.Creat(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	accRes := AccountResponse{
		AccountID:   newAcc.AccountID,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &accRes, nil
}
func (s accountService) GetAccount(customerID int) ([]AccountResponse, error) {
	accounts, err := s.ar.GetAll(customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	accResS := []AccountResponse{}
	for _, account := range accounts {
		accRes := AccountResponse{
			AccountID:   account.AccountID,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		}
		accResS = append(accResS, accRes)
	}
	return accResS, nil
}
