package poeapimodel

// StashTab is exported
type StashTab struct {
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	ID                string `json:"id"`
	Stash             string `json:"stash"`
	StashType         string `json:"stashType"`
	Items             []Item `json:"items"`
	Public            bool   `json:"public"`
	League            string `json:"league"`
}
