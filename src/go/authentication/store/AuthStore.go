package authstore

import "net/http"

// AuthStore is exported
type AuthStore interface {
  Get(req *http.Request) (string, error)
  Set(res http.ResponseWriter, value string)
  Unset(res http.ResponseWriter)
}
