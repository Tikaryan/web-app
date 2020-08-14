package controllers

import (
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func Session(sesName string) *http.Cookie {
	sID, _ := uuid.NewV4()
	// SessionData["uuid"] = sID.String()
	// hhtp.Cookie will give pointer to the Cookie address
	cook := &http.Cookie{
		Name:  sesName,
		Value: sID.String(),
	}
	return cook
}

func SessionValue(key, value string, cook *http.Cookie) (*http.Cookie, error) {
	if key != "" && value != "" {
		cook.Value = cook.Value + "|" + key + "=" + value
		return cook, nil
	}
	return nil, errors.New("empty key or value")
}
