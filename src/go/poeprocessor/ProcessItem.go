package poeprocessor

import (
	poeormmodel "github.com/AdityaHegde/PathOfExileTrade/model/orm/poe"
	"github.com/AdityaHegde/PathOfExileTrade/model/poeapimodel"
	"gorm.io/gorm"
)

// ProcessItems is exported
func ProcessItems(
	db *gorm.DB, account *poeormmodel.Account, apiItems []poeapimodel.Item,
) ([]*poeormmodel.Item, error) {
	var items = make([]*poeormmodel.Item, len(apiItems))

	for i, apiItem := range apiItems {
		items[i], _ = processItem(db, account, apiItem)
	}

	return items, nil
}

func processItem(
	db *gorm.DB, account *poeormmodel.Account, apiItem poeapimodel.Item,
) (*poeormmodel.Item, error) {
	var item poeormmodel.Item = poeormmodel.Item{
		ID:        apiItem.ID,
		Name:      apiItem.Name,
		TypeLine:  apiItem.TypeLine,
		DescrText: apiItem.DescrText,
		League:    apiItem.League,
		X:         apiItem.X,
		Y:         apiItem.Y,
	}

	// TODO: pass on IsImplicit, IsExplicit etc
	if apiItem.ImplicitMods != nil {
		ProcessProperties(db, item, apiItem.ImplicitMods)
	}
	if apiItem.ExplicitMods != nil {
		ProcessProperties(db, item, apiItem.ExplicitMods)
	}
	if apiItem.EnchantMods != nil {
		ProcessProperties(db, item, apiItem.EnchantMods)
	}
	if apiItem.CraftedMods != nil {
		ProcessProperties(db, item, apiItem.CraftedMods)
	}

	return &item, nil
}
