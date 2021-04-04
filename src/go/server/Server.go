package server

import (
  "github.com/AdityaHegde/PathOfExileTrade/account"
  "github.com/AdityaHegde/PathOfExileTrade/authentication"
  "github.com/gorilla/mux"
  "net/http"
)

type Server struct {
  RegisterMiddleware account.RegisterMiddleware
  AuthMiddleware     authentication.AuthMiddleware
}

type SampleResp struct {
  UserName string `json:"name"`
  Sample string `json:"sample"`
}

func (server *Server) Init(handler *mux.Router) error {
  authInit := server.AuthMiddleware.Init()

  if authInit != nil {
    return authInit
  }

  server.setupAuthRoutes(handler)

  server.sampleRestrictedAPI(handler)

  handler.PathPrefix("/").Handler(http.StripPrefix("/",
    http.FileServer(http.Dir("/Users/adityahegde/Git/PathOfExileTrade/public"))))

  return nil
}

func (server *Server) setupAuthRoutes(handler *mux.Router) {
  authRoute := handler.PathPrefix("/auth").Subrouter()

  loginRoute := authRoute.Path("/login")
  loginRoute.Methods(http.MethodPost)
  loginRoute.Handler(server.AuthMiddleware.Login(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    // TODO: remove password from response
    ModelJsonResp(res, req, 200, account.UserContextKey)
  })))

  signupRoute := authRoute.Path("/signup")
  signupRoute.Methods(http.MethodPost)
  signupRoute.Handler(server.RegisterMiddleware.Middleware(server.AuthMiddleware.Signup(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    ModelJsonResp(res, req, 200, account.UserContextKey)
  }))))

  logoutRoute := authRoute.Path("/logout")
  logoutRoute.Methods(http.MethodGet)
  logoutRoute.Handler(server.AuthMiddleware.Logout(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    RawJsonResp(res, 200, new(interface{}))
  })))

  userRoute := authRoute.Path("/user")
  userRoute.Methods(http.MethodGet)
  userRoute.Handler(server.AuthMiddleware.Validate(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    ModelJsonResp(res, req, 200, account.UserContextKey)
  })))
}

func (server *Server) sampleRestrictedAPI(handler *mux.Router) {
  restrictRoute := handler.PathPrefix("/api").Subrouter()
  restrictRoute.Use(server.AuthMiddleware.Validate)

  restrictRoute.HandleFunc("/sample", func(res http.ResponseWriter, req *http.Request) {
    user := req.Context().Value(account.UserContextKey).(account.User)

    resp := SampleResp{
      UserName: user.Name,
      Sample: "Sample Resp",
    }

    RawJsonResp(res, 200, resp)
  })
}
