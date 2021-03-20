package poeprocessor

import (
	"errors"

	"github.com/AdityaHegde/PathOfExileTrade/database"
	poeormmodel "github.com/AdityaHegde/PathOfExileTrade/model/orm/poe"
	"gorm.io/gorm"
)

// ProcessAccount is exported
func ProcessAccount(
	db *gorm.DB, accountName string, lastCharacterName string,
) (*poeormmodel.Account, error) {
	if len(accountName) == 0 {
		return nil, errors.New("Empty account name")
	}
	var account *poeormmodel.Account = &poeormmodel.Account{
		AccountName:       accountName,
		LastCharacterName: lastCharacterName,
	}
	error := database.FindOrCreate(db, account)

	return account, error
}
