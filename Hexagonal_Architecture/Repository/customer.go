package repository

import "fmt"

type Customer struct {
	CustomerID  int    `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	DateOfBirth string `db:"date_of_birth"`
	Gender      string `db:"gender"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}

func Hello() {
	fmt.Println("hello from repository")
}
