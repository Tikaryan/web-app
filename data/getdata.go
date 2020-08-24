package data

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Tikaryan/web-app/config"
	"github.com/Tikaryan/web-app/models"
)

func GetAllUsers() []models.User {
	rows, err := config.DB.Query("SELECT * FROM go_users")
	if err != nil {
		panic(err)
	}
	fmt.Println("*****getAllquery******")
	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		// usr := new(models.User)
		usr := models.User{}
		err = rows.Scan(&usr.ID, &usr.Email, &usr.FirstName, &usr.LastName, &usr.City, &usr.State, &usr.Zipcode, &usr.Password) // order matters
		if err != nil {
			panic(err)
		}
		// scan will insert data in usr bcz we are using "&"
		users = append(users, usr)
	}
	return users
}
func SaveUser(req *http.Request) (string, error) {
	newUser := models.User{}
	newUser.FirstName = req.FormValue("firstName")
	newUser.LastName = req.FormValue("lastName")
	newUser.City = req.FormValue("city")
	newUser.State = req.FormValue("state")
	newUser.Zipcode = req.FormValue("zipcode")
	newUser.Email = req.FormValue("email")
	newUser.Password = []byte(req.FormValue("password"))

	if newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" || newUser.Password == nil {
		return "", errors.New("400. Bad request, All fields must be complete")

	}
	_, err := config.DB.Exec("INSERT INTO go_users (firstName,lastName,city,state,zipcode,email,password) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		newUser.FirstName, newUser.LastName, newUser.City, newUser.State, newUser.Zipcode, newUser.Email, newUser.Password)

	if err != nil {
		return "", errors.New("500. Internal Server Error." + err.Error())
	}
	msg := newUser.Email + " is saved."
	return msg, nil
}

func LoginUser(res http.ResponseWriter, req *http.Request) (*models.User, error) {
	userID := req.FormValue("loginid")
	pass := req.FormValue("password")
	fmt.Println("*****LoginUser******")
	fmt.Println(userID, pass)
	row := config.DB.QueryRow("Select * from go_users where email = $1 AND password = $2", userID, pass)
	usr := models.User{}
	fmt.Println("*****row******")
	fmt.Println(row)
	err := row.Scan(&usr.ID, &usr.Email, &usr.FirstName, &usr.LastName, &usr.City, &usr.State, &usr.Zipcode, &usr.Password) // order matters
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New(http.StatusText(500))
	case err != nil:
		return nil, errors.New(http.StatusText(404))
	}
	fmt.Println("*****getuser******")
	fmt.Println("user= ", usr.Password, usr.FirstName)
	return &usr, nil
}
