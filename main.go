package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73:3306)/banking")
	if err != nil {
		panic(err)
	}

	_ = db
	// customerRepository := repository.NewCustomerRepositoryDB(db)
	// customerService := service.NewCustomerService(customerRepository)

	// customers, err := customerService.GetCustomers()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customers)

	// customer, err := customerService.GetCustomer(2000)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)
}
