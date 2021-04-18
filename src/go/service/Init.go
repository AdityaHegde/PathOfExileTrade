package service

import "gorm.io/gorm"

// Init is exported
func Init(db* gorm.DB) {
  db.AutoMigrate(ListingType{})
  db.AutoMigrate(Listing{})
  db.AutoMigrate(Fulfillment{})
}
