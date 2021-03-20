package service

import (
	"time"

	accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"
	servicemodel "github.com/AdityaHegde/PathOfExileTrade/model/service"
	"gorm.io/gorm"
)

// CreateListing is exported
func CreateListing(
	db *gorm.DB, User accountmodel.User,
	listingType string, maxParticipants uint, count uint,
) (*servicemodel.Listing, error) {
	var listing = servicemodel.Listing{
		User:            User,
		ListingTypeID:   listingType,
		MaxParticipants: maxParticipants,
		Count:           count,
		Active:          false,
	}

	res := db.Create(&listing)

	return &listing, res.Error
}

// UpdateListing is exported
func UpdateListing(
	db *gorm.DB, listingID uint64, updatedListing servicemodel.Listing,
) (*servicemodel.Listing, error) {
	var listing = servicemodel.Listing{
		ID: listingID,
	}

	if !updatedListing.Active {
		updatedListing.LastActiveAt = time.Now()
	}

	res := db.Model(&listing).Updates(updatedListing)

	return &listing, res.Error
}
