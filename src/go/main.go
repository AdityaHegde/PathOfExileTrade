package main

import (
	"fmt"
	"github.com/AdityaHegde/PathOfExileTrade/account"
	"github.com/AdityaHegde/PathOfExileTrade/authentication"
	authcore "github.com/AdityaHegde/PathOfExileTrade/authentication/core"
	authstore "github.com/AdityaHegde/PathOfExileTrade/authentication/store"
	"github.com/AdityaHegde/PathOfExileTrade/database"
	"github.com/AdityaHegde/PathOfExileTrade/service"
	"net/http"

	"github.com/AdityaHegde/PathOfExileTrade/server"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("DB connected")
		account.Init(db)
		service.Init(db)
	}

	serverInst := server.Server{
		RegisterMiddleware: account.RegisterMiddleware{
			Db: db,
		},
		AuthMiddleware: authentication.AuthMiddleware{
			Auth: new(authcore.JwtAuth),
			Db: db,
			Store: new(authstore.CookieStore),
		},
	}

	router := mux.NewRouter()
	initErr := serverInst.Init(router)
	if initErr != nil {
		fmt.Println(err)
	} else {
		fmt.Println(http.ListenAndServe(":3000", router))
	}
}
