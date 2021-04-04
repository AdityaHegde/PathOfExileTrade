package authentication

import (
  "fmt"
  "github.com/AdityaHegde/PathOfExileTrade/account"
  authcore "github.com/AdityaHegde/PathOfExileTrade/authentication/core"
  authstore "github.com/AdityaHegde/PathOfExileTrade/authentication/store"
  "github.com/google/jsonapi"
  "gorm.io/gorm"
  "net/http"
)

const UserParam = "user"
const PasswordParam = "pwd"

// AuthMiddleware is exported
type AuthMiddleware struct {
  Store authstore.AuthStore
  Auth authcore.Auth
  Db   *gorm.DB
}

func (authMiddleware *AuthMiddleware) Init() error {
  return authMiddleware.Auth.Init()
}

func (authMiddleware *AuthMiddleware) Validate(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    jwt, err := authMiddleware.Store.Get(req)
    if err != nil {
      fmt.Println(err)
      http.Error(res, "Forbidden", http.StatusForbidden)
      return
    }

    userName, validateErr := authMiddleware.Auth.Validate(jwt)
    if validateErr !=nil {
      http.Error(res, "Forbidden", http.StatusForbidden)
      return
    }

    var user = account.User{
      Name: userName,
    }
    findResp := authMiddleware.Db.Find(&user)

    if findResp.Error != nil || findResp.RowsAffected == 0 {
      fmt.Println(findResp.Error)
      http.Error(res, "Forbidden", http.StatusForbidden)
    } else {
      next.ServeHTTP(res, user.GetRequestWithUser(req))
    }
  })
}

func (authMiddleware *AuthMiddleware) Login(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    var user = new(account.User)
    parseErr := jsonapi.UnmarshalPayload(req.Body, user)
    if parseErr != nil {
      fmt.Println(parseErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }

    passwordFromReq := user.Password

    findResp := authMiddleware.Db.Find(user)
    if findResp.Error != nil || findResp.RowsAffected == 0 {
      fmt.Println(findResp.Error)
      http.Error(res, "Forbidden", http.StatusForbidden)
      return
    }

    checkPwdErr := user.CheckPassword(passwordFromReq)
    if checkPwdErr != nil {
      fmt.Println(checkPwdErr)
      http.Error(res, "Forbidden", http.StatusForbidden)
      return
    }

    authValue, authErr := authMiddleware.Auth.Generate(user)
    if authErr != nil {
      fmt.Println(authErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    authMiddleware.Store.Set(res, authValue)

    next.ServeHTTP(res, user.GetRequestWithUser(req))
  })
}

func (authMiddleware *AuthMiddleware) Signup(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    user := req.Context().Value(account.UserContextKey).(account.User)

    authValue, authErr := authMiddleware.Auth.Generate(&user)
    if authErr != nil {
      fmt.Println(authErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    authMiddleware.Store.Set(res, authValue)

    next.ServeHTTP(res, user.GetRequestWithUser(req))
  })
}

func (authMiddleware *AuthMiddleware) Logout(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    authMiddleware.Store.Unset(res)
    next.ServeHTTP(res, req)
  })
}

func (authMiddleware *AuthMiddleware) Restrict(next http.Handler, restrictedRole account.UserRole) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    user := req.Context().Value(account.UserContextKey).(account.User)

    if user.Role <= restrictedRole {
      next.ServeHTTP(res, req)
    } else {
      http.Error(res, "Forbidden", http.StatusForbidden)
    }
  })
}
