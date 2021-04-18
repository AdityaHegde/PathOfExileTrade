package account

import "gorm.io/gorm"

// Init is exported
func Init(db *gorm.DB) {
  db.AutoMigrate(User{})
}
