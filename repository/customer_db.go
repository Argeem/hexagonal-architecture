package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct { // start with : lower case private : upper case public on the other package
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) { // เป็นเหมือนการ implement ใน java
	customers := []Customer{}
	query := `select * from customers`
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := `select * from customers where id=?`
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
