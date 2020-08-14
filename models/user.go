package models

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	City      string
	State     string
	Zipcode   string
	Password  []byte
}
