package service

import (
  "github.com/AdityaHegde/PathOfExileTrade/account"
  "github.com/AdityaHegde/PathOfExileTrade/authentication"
  "github.com/gorilla/mux"
  "gorm.io/gorm"
  "net/http"
)

// Routes is exported
type Routes struct {
  db             *gorm.DB
  AuthMiddleware authentication.AuthMiddleware
}

func (serviceRoutes *Routes) Init(router *mux.Router) {
  serviceRouter := router.PathPrefix("/api/service").Subrouter()
  serviceRouter.Use(serviceRoutes.AuthMiddleware.Validate)

  serviceRoutes.registerListingType(serviceRouter)
}

func (serviceRoutes *Routes) registerListingType(router *mux.Router) {
  listingTypeRouter := router.PathPrefix("/listingTypes").Subrouter()
  listingTypeRouter.Use(serviceRoutes.AuthMiddleware.Restrict(account.Owner))

  listingTypeRouter.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    HandleGetAllListingType(res, serviceRoutes.db)
  })
  listingTypeRouter.HandleFunc("/{id}", func(res http.ResponseWriter, req *http.Request) {
    HandleIDListingType(res, req, serviceRoutes.db)
  })
}
