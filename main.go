package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()
	db := initDatabase()
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryDB
	_ = customerRepositoryMock
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID}", customerHandler.GetCustomer).Methods(http.MethodPost)

	http.ListenAndServe(fmt.Sprintf("%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()                                   // check environment if exist use it instead
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // change environment

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok") // base on time zone database
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.databasr"),
	)
	driver := viper.GetString("app.port")
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute) // set expire time
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
