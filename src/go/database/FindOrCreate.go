package database

import "gorm.io/gorm"

// FindOrCreate is imported
func FindOrCreate(db *gorm.DB, model interface{}) error {
	result := db.First(model)

	if result.RowsAffected == 1 {
		return nil
	}

	result = db.Create(model)

	if result.Error != nil {
		return nil
	}

	return result.Error
}
