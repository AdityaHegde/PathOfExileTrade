package main

import (
	"fmt"

	"github.com/AdityaHegde/PathOfExileTrade/client"
	"github.com/AdityaHegde/PathOfExileTrade/database"
	"github.com/AdityaHegde/PathOfExileTrade/poeprocessor"
)

func main() {
	db, err := database.Connect()
	// _, err := database.Connect()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB connected")
	}

	stashTabs, err := client.GetNextPublicStashTabs(
		"1103640760-1111112799-1071388634-1201035544-1152442178")

	poeprocessor.ProcessStashTabs(db, *stashTabs)
}
