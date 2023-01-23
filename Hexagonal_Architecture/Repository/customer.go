package Repository

type Customer struct {
	CustomerID  int    `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Dateofbirth string `db:"date_of_birth"`
	Gender      string `db:"gender"`
}

type CustomerRepository interface {
	GetAll()
	GetByID(int)
}
