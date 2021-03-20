package servicemodel

import accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"

// Fulfillment is exported
type Fulfillment struct {
	ListingID uint64
	Listing   Listing
	UserID    string `gorm:"references:Name"`
	User      accountmodel.User
}
