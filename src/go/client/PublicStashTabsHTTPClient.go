package client

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdityaHegde/PathOfExileTrade/model/poeapimodel"
)

const apiURL = "http://api.pathofexile.com/public-stash-tabs/?id="

func getPublicStashTabs(url string) (*poeapimodel.PublicStashTabs, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var respModel poeapimodel.PublicStashTabs

	if err := json.NewDecoder(resp.Body).Decode(&respModel); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &respModel, nil
}

// GetNextPublicStashTabs is exported
func GetNextPublicStashTabs(nextID string) (*poeapimodel.PublicStashTabs, error) {
	return getPublicStashTabs(apiURL + nextID)
}
