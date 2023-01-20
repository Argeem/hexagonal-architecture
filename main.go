package main

import (
	"bank/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73:3306)/banking")
	if err == nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	_ = customerRepository
}
