package service

import (
	servicemodel "github.com/AdityaHegde/PathOfExileTrade/model/service"
	"gorm.io/gorm"
)

// CreateListingType is exported
func CreateListingType(db *gorm.DB, listingTypeName string) error {
	var listingType = servicemodel.ListingType{
		Name:   listingTypeName,
		Active: true,
	}

	res := db.Create(&listingType)

	return res.Error
}

// UpdateListingType is exported
func UpdateListingType(
	db *gorm.DB, listingType servicemodel.ListingType,
) (*servicemodel.ListingType, error) {
	res := db.Model(&listingType).Updates(listingType)

	return &listingType, res.Error
}

// GetAllListingType is exported
func GetAllListingType(db *gorm.DB) (*[]servicemodel.ListingType, error) {
	var listingTypes []servicemodel.ListingType

	res := db.Find(&listingTypes)

	if res.Error != nil {
		return nil, res.Error
	}

	return &listingTypes, nil
}
