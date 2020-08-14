package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	City      string
	State     string
	Zipcode   string
	Email     string
	Password  []byte
}
