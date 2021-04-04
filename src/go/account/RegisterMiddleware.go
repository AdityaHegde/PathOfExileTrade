package account

import (
  "fmt"
  "github.com/google/jsonapi"
  "gorm.io/gorm"
  "net/http"
)

// RegisterMiddleware is exported
type RegisterMiddleware struct {
  Db *gorm.DB
}

func (register *RegisterMiddleware) Middleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    var user = new(User)
    parseErr := jsonapi.UnmarshalPayload(req.Body, user)
    if parseErr != nil {
      fmt.Println(parseErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    user.Role = Consumer

    existing := register.Db.Find(user)

    if existing.RowsAffected > 0 {
      fmt.Println("Username already exists")
      // TODO: return readable error for username exists
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }

    hashErr := user.HashPassword(user.Password)
    if hashErr != nil {
      fmt.Println(hashErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }

    createErr := user.CreateUserRecord(register.Db)
    if createErr != nil {
      fmt.Println(createErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
  })
}
