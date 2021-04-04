package authstore

import (
	"net/http"
	"time"
)

const CookieName = "jwt_auth"

// CookieStore is exported
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
	setCookie(res, value, time.Now().Add(120 * time.Minute))
}

func (cookie *CookieStore) Unset(res http.ResponseWriter) {
	setCookie(res, "", time.Now().Add(-120 * time.Minute))
}

func setCookie(res http.ResponseWriter, value string, expiry time.Time) {
	http.SetCookie(res, &http.Cookie{
		Name:    CookieName,
		Value:   value,
		Expires: expiry,
		Domain:  "",
		Path:    "/",
	})
}
