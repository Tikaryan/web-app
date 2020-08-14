package controllers

import (
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var SessionData = make(map[string]string)

func Session(sessName string) *http.Cookie {
	sID, _ := uuid.NewV4()
	SessionData["uuid"] = sID.String()
	// hhtp.Cookie will give pointer to the Cookie address
	cook := &http.Cookie{
		Name:  sessName,
		Value: SessionData["uuid"],
	}
	return cook
}

func SessionValue(key, value string) (map[string]string, error) {
	if key != "" && value != "" {
		SessionData[key] = value
		return SessionData, nil
	}
	return nil, errors.New("empty key or value")
}
