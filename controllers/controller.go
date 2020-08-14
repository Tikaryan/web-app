package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Tikaryan/web-app/data"
	"github.com/Tikaryan/web-app/models"
)

var tpl *template.Template

type allUsers []models.User

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func LoginPage(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "loginPage.html", nil)
	if err != nil {
		log.Fatalln("error Opening login page", http.StatusNotFound)
	}
}
func GetUser(res http.ResponseWriter, req *http.Request) {

}
func CheckUser(res http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("loginid")
	var checkUser bool
	users := data.GetAllUsers()
	for _, v := range users {
		if userID == v.Email {
			checkUser = true
			break
		}
	}
	res.Header().Set("Content-type", "application/json")
	fmt.Fprintln(res, checkUser)
}
func SignupPage(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "signupPage.html", nil)
	if err != nil {
		log.Fatalln("error Opening login page", http.StatusNotFound)
	}
}
func CreateUser(res http.ResponseWriter, req *http.Request) {
	sessionData := make(map[string]string)

	if req.Method == http.MethodPost {
		// login part start
		session := Session("session")
		http.SetCookie(res, session)
		// login part end
		msg, err := data.SaveUser(req)
		if err != nil {
			sessionData, _ = SessionValue("error", err.Error())
		} else {

			sessionData, _ = SessionValue("msg", msg)
		}
		tpl.ExecuteTemplate(res, "loginPage.html", sessionData)
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	cook, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	cook = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	for k := range SessionData {
		delete(SessionData, k)
	}
	http.SetCookie(res, cook)
}
