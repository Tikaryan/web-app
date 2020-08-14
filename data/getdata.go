package data

import (
	"errors"
	"net/http"

	"github.com/Tikaryan/web-app/config"
	"github.com/Tikaryan/web-app/models"
)

func GetAllUsers() []models.User {
	rows, err := config.DB.Query("Select * from go_users")
	if err != nil {
		panic(err)
	}
	rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		// usr := new(models.User)
		usr := models.User{}
		err = rows.Scan(&usr.ID, &usr.FirstName, &usr.LastName, &usr.City, &usr.State, &usr.Email, &usr.Zipcode, &usr.Password) // order matters
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
		return "", errors.New("400. Bad request. All fields must be complete")

	}
	_, err := config.DB.Exec("INSERT INTO go_users (firstName,lastName,city,state,zipcode,email,password) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		newUser.FirstName, newUser.LastName, newUser.City, newUser.State, newUser.Zipcode, newUser.Email, newUser.Password)

	if err != nil {
		return "", errors.New("500. Internal Server Error." + err.Error())
	}
	msg := "Success" + newUser.Email + "is saved"
	return msg, nil
}
