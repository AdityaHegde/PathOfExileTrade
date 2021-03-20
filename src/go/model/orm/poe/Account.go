package poeormmodel

// Account is exported
type Account struct {
	AccountName       string `gorm:"primaryKey"`
	LastCharacterName string
}
