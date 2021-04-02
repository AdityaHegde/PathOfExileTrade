package authstore

import "net/http"

type AuthStore interface {
  Get(req *http.Request) (string, error)
  Set(res http.ResponseWriter, value string)
}
