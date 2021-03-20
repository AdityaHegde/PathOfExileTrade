package poeormmodel

import (
	"time"
)

// Item is exported
type Item struct {
	ID         string     `gorm:"primaryKey"`
	UpdatedAt  time.Time  `gorm:"index"`
	StashTabID string     `gorm:"index"`
	Name       string     `gorm:"index"`
	Categories []Category `gorm:"many2many:item_categories"`
	TypeLine   string
	DescrText  string
	Mods       []PropertyValue `gorm:"foreignKey:ItemID"`

	PriceDenomination string
	PriceValue        float32
	IsExactPrice      bool

	X      uint
	Y      uint
	Icon   string
	League string
	Note   string
}
