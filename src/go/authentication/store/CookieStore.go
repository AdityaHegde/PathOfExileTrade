package authstore

import (
	"net/http"
	"time"
)

const CookieName = "jwt_auth"

type CookieStore struct {
}

func (cookie *CookieStore) Get(req *http.Request) (string, error) {
	cookieValue, err := req.Cookie(CookieName)

	if err != nil {
		return "", err
	} else {
		return cookieValue.Value, nil
	}
}

func (cookie *CookieStore) Set(res http.ResponseWriter, value string) {
	http.SetCookie(res, &http.Cookie{
		Name:    CookieName,
		Value:   value,
		Expires: time.Now().Add(120 * time.Minute),
		Domain:  "",
		Path:    "/",
	})
}
