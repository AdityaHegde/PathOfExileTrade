package poeormmodel

// StashTab is exported
type StashTab struct {
	ID          string  `gorm:"primaryKey"`
	AccountName string  `gorm:"index"`
	Account     Account `gorm:"foreignKey:AccountName;references:AccountName"`
	Stash       string
	StashType   string
	Items       []Item `gorm:"foreignKey:StashTabID"`
	Public      bool
	League      string
}
