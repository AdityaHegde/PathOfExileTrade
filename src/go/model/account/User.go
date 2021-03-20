package accountmodel

// User is extended
type User struct {
	Name         string `gorm:"primaryKey"`
	PasswordSalt string
}
