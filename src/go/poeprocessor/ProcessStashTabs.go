package poeprocessor

import (
	"github.com/AdityaHegde/PathOfExileTrade/database"
	poeormmodel "github.com/AdityaHegde/PathOfExileTrade/model/orm/poe"
	"github.com/AdityaHegde/PathOfExileTrade/model/poeapimodel"
	"gorm.io/gorm"
)

// ProcessStashTabs is exported
func ProcessStashTabs(db *gorm.DB, stashTabs poeapimodel.PublicStashTabs) {
	for _, apiStashTab := range stashTabs.StashTabs {
		_, err := processStashTab(db, apiStashTab)

		if err != nil {
			continue
		}

		account, err := ProcessAccount(db, apiStashTab.AccountName, apiStashTab.LastCharacterName)

		if err == nil {
			ProcessItems(db, account, apiStashTab.Items)
		}
		// TODO: delete any items not in the stash anymore
		// make sure items moved to other stashes are not deleted and added again
	}
}

// Defines how poeapimodel.StashTab maps to poeormmodel.StashTab
func processStashTab(db *gorm.DB, apiStashTab poeapimodel.StashTab) (*poeormmodel.StashTab, error) {
	var stashTab *poeormmodel.StashTab = &poeormmodel.StashTab{
		ID:          apiStashTab.ID,
		AccountName: apiStashTab.AccountName,
		Stash:       apiStashTab.Stash,
		StashType:   apiStashTab.StashType,
		Public:      apiStashTab.Public,
		League:      apiStashTab.League,
	}
	error := database.FindOrCreate(db, stashTab)
	return stashTab, error
}
