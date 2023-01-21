package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73:3306)/banking")
	if err != nil {
		panic(err)
	}
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	_ = customerRepositoryDB
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID}", customerHandler.GetCustomer).Methods(http.MethodPost)

	http.ListenAndServe(":8000", router)
}
