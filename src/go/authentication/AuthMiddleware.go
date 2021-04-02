package authentication

import (
  "context"
  authcore "github.com/AdityaHegde/PathOfExileTrade/authentication/core"
  authstore "github.com/AdityaHegde/PathOfExileTrade/authentication/store"
  accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"
  "gorm.io/gorm"
  "net/http"
)

const UserParam = "user"
const PasswordParam = "pwd"
// UserContextKey is exported
const UserContextKey = "user"

// AuthMiddleware is exported
type AuthMiddleware struct {
  store authstore.AuthStore
  auth authcore.Auth
  db *gorm.DB
}

func (authMiddleware *AuthMiddleware) Validate(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    userName, err := authMiddleware.store.Get(req)
    if err != nil {
      http.Error(res, "Forbidden", http.StatusForbidden)
    } else {
      var user = accountmodel.User{
        Name: userName,
      }
      findResp := authMiddleware.db.Find(&user)

      if findResp.Error != nil || findResp.RowsAffected == 0 {
        http.Error(res, "Forbidden", http.StatusForbidden)
      } else {
        next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), UserContextKey, user)))
      }
    }
  })
}

func (authMiddleware *AuthMiddleware) Login(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    var user = &accountmodel.User{
      Name: req.URL.Query()[UserParam][0],
    }
    findResp := authMiddleware.db.Find(user)

    if findResp.Error != nil || findResp.RowsAffected == 0 {
      http.Error(res, "Forbidden", http.StatusForbidden)
      return
    }

    checkPwdErr := user.CheckPassword(req.URL.Query()[PasswordParam][0])
    if checkPwdErr != nil {
      http.Error(res, "Forbidden", http.StatusForbidden)
      return
    }

    next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), UserContextKey, user)))
  })
}

func (authMiddleware *AuthMiddleware) Signup(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    var user = &accountmodel.User{
      Name: req.URL.Query()[UserParam][0],
      Roles: []accountmodel.UserRole{accountmodel.Consumer},
    }
    hashErr := user.HashPassword(req.URL.Query()[PasswordParam][0])

    if hashErr.Error != nil {
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }

    authValue, authErr := authMiddleware.auth.Generate(user)
    if authErr != nil {
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    authMiddleware.store.Set(res, authValue)

    createErr := user.CreateUserRecord(authMiddleware.db)
    if createErr != nil {
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }

    next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), "user", user)))
  })
}

func (authMiddleware *AuthMiddleware) Restrict(next http.Handler, restrictedRole accountmodel.UserRole) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    user := req.Context().Value(UserContextKey).(accountmodel.User)

    for role := range user.Roles {
      if role == int(restrictedRole) {
        next.ServeHTTP(res, req)
        return
      }
    }

    http.Error(res, "Forbidden", http.StatusForbidden)
  })
}
