package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct { // start with : lower case private : upper case public on the other package
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}
