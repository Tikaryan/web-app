package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
)

var store = sessions.NewCookieStore([]byte("secrete-passkey"))

// var store *sessions
func Session(req *http.Request, res http.ResponseWriter, sessName string) *sessions.Session {
	session, _ := store.Get(req, sessName)
	err := session.Save(req, res)
	if err != nil {
		log.Fatalln(err)
	}
	return session
}

func Cookies(sesName string) *http.Cookie {
	sID, _ := uuid.NewV4()
	// SessionData["uuid"] = sID.String()
	// hhtp.Cookie will give pointer to the Cookie address
	cook := &http.Cookie{
		Name:  sesName,
		Value: sID.String(),
	}
	fmt.Println("*****session******")
	fmt.Println(cook)
	return cook
}

func CookiesValue(key, value string, cook *http.Cookie) (*http.Cookie, error) {
	if key != "" && value != "" {
		cook.Value = cook.Value + "|" + key + "=" + value
		return cook, nil
	}
	fmt.Println("*****cookiesValue******")
	fmt.Println(cook)
	return nil, errors.New("empty key or value")
}
