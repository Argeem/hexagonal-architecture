package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Creat(acc Account) (*Account, error) {
	query := `insert into accounts (customer_id, opening_date, account_type, amount, status ) value(?, ?, ?, ?, ?)`
	result, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	acc.AccountID = int(id)
	return &acc, nil
}
func (r accountRepositoryDB) GetAll(CustomerID int) ([]Account, error) {
	query := `select * form accounts where customer_id=?`
	account := []Account{}
	err := r.db.Select(&account, query, CustomerID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
