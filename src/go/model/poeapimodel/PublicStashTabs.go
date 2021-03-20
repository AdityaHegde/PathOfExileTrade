package poeapimodel

// PublicStashTabs is exported
type PublicStashTabs struct {
	NextChangeID string     `json:"next_change_id"`
	StashTabs    []StashTab `json:"stashes"`
}
