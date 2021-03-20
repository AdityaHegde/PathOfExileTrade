package servicemodel

import (
	"time"

	accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"
)

// Listing is exported
type Listing struct {
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	UserID          string
	User            accountmodel.User `gorm:"references:Name"`
	ListingTypeID   string
	ListingType     ListingType `gorm:"references:Name"`
	Active          bool        `gorm:"index"`
	MaxParticipants uint        `gorm:"index"`
	Count           uint        `gorm:"index"`
	LastActiveAt    time.Time   `gorm:"index"`
	CreatedAt       time.Time   `gorm:"autoUpdateTime:milli"`
	UpdatedAt       time.Time   `gorm:"autoCreateTime"`
}
