package service

import (
	"github.com/AdityaHegde/PathOfExileTrade/account"
)

// Fulfillment is exported
type Fulfillment struct {
	ListingID uint64
	Listing   Listing
	UserID    string `gorm:"references:Name"`
	User      account.User
}
