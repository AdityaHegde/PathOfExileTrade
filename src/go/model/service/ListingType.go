package servicemodel

// ListingType is exported
type ListingType struct {
	Name   string `gorm:"primaryKey"`
	Active bool   `gorm:"index"`
}
