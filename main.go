package main

import (
	"net/http"

	controller "github.com/Tikaryan/web-app/controllers"
	"github.com/gorilla/context"
)

func main() {
	cssHand := http.FileServer(http.Dir("./templates/resources/css/"))
	imgHand := http.FileServer(http.Dir("./templates/resources/images/"))
	jsHand := http.FileServer(http.Dir("./templates/resources/javascript/"))

	http.HandleFunc("/loginpage", controller.LoginPage)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/loginAuth", controller.LoginAuth)
	http.HandleFunc("/checkUser", controller.CheckUser)
	http.HandleFunc("/signup", controller.SignupPage)
	http.HandleFunc("/createUser", controller.CreateUser)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/dashboard", controller.Dashboard)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.Handle("/resources/css/", http.StripPrefix("/resources/css", cssHand))
	http.Handle("/resources/images/", http.StripPrefix("/resources/images", imgHand))
	http.Handle("/resources/javascript/", http.StripPrefix("/resources/javascript", jsHand))

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))

	// http.Handle("/resources/", r)
}

// func (r resHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	url := req.RequestURI
// 	s := strings.Split(url, "/")[1]
// 	if s == "images" {
// 		http.FileServer(http.Dir("./templates/resources/images/"))
// 	} else if s == "css" {
// 		http.FileServer(http.Dir("./templates/resources/css/"))
// 	}
// }
