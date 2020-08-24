package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Tikaryan/web-app/data"
	"github.com/Tikaryan/web-app/models"
)

var tpl *template.Template
var tpl1 *template.Template

type allUsers []models.User

//very important point to remember for composite literals,
// During inserting data we have to tell what type of data is coming like userData{ array: []TYPEOFARRAY(i.e string,int,..){"Then initialize it"} }
type UserData struct {
	// we are defing that userData struct will have a map inside it, we have not initialized,
	// we have to initialize it before entering the data
	Maps map[string]interface{} `json:"Data"`
}

var JsonData = make(map[string]interface{})

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func LoginPage(res http.ResponseWriter, req *http.Request) {
	session := Session(req, res, "session")
	session.Options.MaxAge = -1
	cook := &http.Cookie{
		Name:   "temp-cookie",
		MaxAge: -1,
	}
	http.SetCookie(res, cook)
	session.Save(req, res)
	err := tpl.ExecuteTemplate(res, "loginPage.html", nil)
	if err != nil {
		log.Fatalln("error Opening login page", http.StatusNotFound)
	}
}
func LoginAuth(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		res.Header().Set("Content-type", "application/json")
		var marData = []byte{}
		// initializing map using make so we can enter data and convert it into json to send

		usr, err := data.LoginUser(res, req)
		if err != nil {
			fmt.Println("*****Error******")
			log.Println(err, err.Error())
			JsonData["error"] = err.Error()
			m := UserData{
				Maps: JsonData,
			}
			marData, _ = json.Marshal(m)
			delete(JsonData, "error")
			res.Header().Set("Content-type", "application/json")
			res.Write(marData)
		} else {
			session := Session(req, res, "session")
			session.Values["email"] = usr.Email
			JsonData["success"] = "success"
			m := UserData{
				Maps: JsonData,
			}
			marData, _ = json.Marshal(m)
			delete(JsonData, "success")
			err = session.Save(req, res)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-type", "application/json")
			res.Write(marData)
		}
		// tpl.ExecuteTemplate(res, "dashboard.html", nil)
	} else {
		http.Error(res, http.StatusText(400), 400)
	}

}
func Login(res http.ResponseWriter, req *http.Request) {
	session := Session(req, res, "session")
	if req.FormValue("email") == session.Values["email"] && req.Method == http.MethodPost {
		tpl.ExecuteTemplate(res, "dashboard.html", nil)
	} else {
		http.Error(res, http.StatusText(406), http.StatusNotAcceptable)
	}
}
func CheckUser(res http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("loginid")
	var checkUser bool
	users := data.GetAllUsers()
	fmt.Println("*****getUser******")
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
		log.Fatalln("error Opening Signup page", http.StatusNotFound)
	}
}
func CreateUser(res http.ResponseWriter, req *http.Request) {
	cookie := Cookies("temp-cookie")
	if req.Method == http.MethodPost {
		msg, err := data.SaveUser(req)
		if err != nil {
			cookie, _ = CookiesValue("error", err.Error(), cookie)
		} else {
			cookie, _ = CookiesValue("msg", msg, cookie)
		}
		http.SetCookie(res, cookie)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.ExecuteTemplate(res, "loginPage.html", nil)
		cookie.MaxAge = -1
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	cook, err := req.Cookie("session")
	session := Session(req, res, "session")
	session.Options.MaxAge = -1
	session.Save(req, res)
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	cook = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, cook)

	// redirect not working
	// http.Redirect(res, req, "loginpage", http.StatusSeeOther)

	// this is also not working for now..need to figure out why?

	err = tpl.ExecuteTemplate(res, "loginPage.html", nil)
	if err != nil {
		http.Error(res, http.StatusText(404), http.StatusBadRequest)
	}
}

func Dashboard(res http.ResponseWriter, req *http.Request) {
	tpl1 = template.Must(template.ParseGlob("templates/gohtmls/*.gohtml"))
	JsonData["Error"] = "error"
	m := UserData{
		Maps: JsonData,
	}
	err := tpl.ExecuteTemplate(res, "dashboard.html", m)
	if err != nil {
		log.Fatalln("error Dashboard login page", http.StatusNotFound)
	}
}
