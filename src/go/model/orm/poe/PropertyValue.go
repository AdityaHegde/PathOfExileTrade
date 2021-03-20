package poeormmodel

// PropertyValue is exported
type PropertyValue struct {
	ItemID           string   `gorm:"primaryKey"`
	PropertyID       string   `gorm:"primaryKey"`
	Property         Property `gorm:"references:PropertyName"`
	PropertyMinValue string   `gorm:"index"`
	PropertyMaxValue string   `gorm:"index"`
}
