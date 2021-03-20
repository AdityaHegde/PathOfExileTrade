package poeormmodel

// Property is exported
type Property struct {
	PropertyName string `gorm:"primaryKey"`

	IsRange  bool
	IsNumber bool
	IsString bool

	IsImplicit bool
	IsExplicit bool
	IsEnchat   bool
	IsCrafted  bool
	IsProperty bool
	IsSuffix   bool
	IsPrefix   bool
}
