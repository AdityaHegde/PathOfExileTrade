package database

import (
	poeormmodel "github.com/AdityaHegde/PathOfExileTrade/model/orm/poe"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect is exported
func Connect() (*gorm.DB, error) {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&poeormmodel.Account{})
	db.AutoMigrate(&poeormmodel.Category{})
	db.AutoMigrate(&poeormmodel.Item{})
	db.AutoMigrate(&poeormmodel.Property{})
	db.AutoMigrate(&poeormmodel.PropertyValue{})
	db.AutoMigrate(&poeormmodel.StashTab{})

	return db, nil
}
