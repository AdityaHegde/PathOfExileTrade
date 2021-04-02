package main

import (
	"fmt"
	"net/http"

	"github.com/AdityaHegde/PathOfExileTrade/server"
	"github.com/gorilla/mux"
)

func main() {
	// db, err := database.Connect()

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("DB connected")
	// }

	router := mux.NewRouter()
	server.SetupAuthentication(router)

	fmt.Println(http.ListenAndServe(":3000", router))
}
