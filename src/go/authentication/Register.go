package authentication

import (
	"errors"

	accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"
	"gorm.io/gorm"
)

// Register is exported
func Register(db *gorm.DB, user accountmodel.User) (*accountmodel.User, error) {
	existing := db.Find(&user)

	if existing.RowsAffected > 0 {
		return nil, errors.New("User name already taken")
	}

	createRes := db.Create(&user)

	return &user, createRes.Error
}
