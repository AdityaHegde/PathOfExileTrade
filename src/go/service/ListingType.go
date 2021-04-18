package service

import (
	"gorm.io/gorm"
)

// ListingType is exported
type ListingType struct {
	Name   string `jsonapi:"primary,users" gorm:"primaryKey"`
	Active bool   `jsonapi:"attr,active" gorm:"index"`
}

// GetListingType is exported
func GetListingType(db *gorm.DB, listingTypeName string) (*ListingType, error) {
	var listingType ListingType

	res := db.Find(&listingType, listingTypeName)

	return &listingType, res.Error
}

// CreateListingType is exported
func CreateListingType(db *gorm.DB, listingType *ListingType) (*ListingType, error) {
	listingType.Active = true

	res := db.Create(listingType)

	return listingType, res.Error
}

// UpdateListingType is exported
func UpdateListingType(db *gorm.DB, listingType *ListingType) (*ListingType, error) {
	res := db.Model(listingType).Updates(listingType)

	return listingType, res.Error
}

// GetAllListingType is exported
func GetAllListingType(db *gorm.DB) (*[]ListingType, error) {
	var listingTypes []ListingType

	res := db.Find(&listingTypes)

	return &listingTypes, res.Error
}

// GetAllActiveListingType is exported
func GetAllActiveListingType(db *gorm.DB) (*[]ListingType, error) {
	var listingTypes []ListingType

	res := db.Find(&listingTypes, "active = ?", true)

	return &listingTypes, res.Error
}
