package service

import (
	"github.com/AdityaHegde/PathOfExileTrade/account"
	"time"

	"gorm.io/gorm"
)

// Listing is exported
type Listing struct {
	ID              uint64 `jsonapi:"primary,users" gorm:"primaryKey;autoIncrement"`
	UserID          string
	User            account.User `jsonapi:"relation,users" gorm:"references:Name"`
	ListingTypeID   string
	ListingType     ListingType `jsonapi:"relation,listingTypes" gorm:"references:Name"`
	Active          bool        `jsonapi:"attr,active" gorm:"index"`
	MaxParticipants uint        `jsonapi:"attr,maxParticipants" gorm:"index"`
	Count           uint        `jsonapi:"attr,count" gorm:"index"`
	LastActiveAt    time.Time   `jsonapi:"attr,lastActiveAt" gorm:"index"`
	CreatedAt       time.Time   `jsonapi:"attr,createdAt" gorm:"autoUpdateTime:milli"`
	UpdatedAt       time.Time   `jsonapi:"attr,updatedAt" gorm:"autoCreateTime"`
}

// CreateListing is exported
func CreateListing(
	db *gorm.DB, User account.User,
	listingType string, maxParticipants uint, count uint,
) (*Listing, error) {
	var listing = Listing{
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
	db *gorm.DB, listingID uint64, updatedListing *Listing,
) (*Listing, error) {
	var listing = Listing{
		ID: listingID,
	}

	if !updatedListing.Active {
		updatedListing.LastActiveAt = time.Now()
	}

	res := db.Model(&listing).Updates(updatedListing)

	return &listing, res.Error
}

// GetOwnListings is exported
func GetOwnListings(db *gorm.DB, user *account.User) (*[]Listing, error) {
	var listings []Listing

	res := db.Find(&listings, "user_id = ?", user.Name)

	return &listings, res.Error
}

// GetActiveListings is exported
func GetActiveListings(db *gorm.DB) (*[]Listing, error) {
	var listings []Listing

	res := db.Find(&listings, "active = ?", true)

	return &listings, res.Error
}
