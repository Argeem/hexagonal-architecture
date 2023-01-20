package repository

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DataOfBirth string `db:"data_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	getAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}