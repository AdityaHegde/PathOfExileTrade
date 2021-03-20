package database

import (
	"sync"

	"gorm.io/gorm"
)

// FindOrCreateInMap is exported
func FindOrCreateInMap(db *gorm.DB, syncMap *sync.Map, key interface{}, model interface{}) error {
	_, loaded := syncMap.LoadOrStore(key, true)

	if !loaded {
		return nil
	}

	return FindOrCreate(db, model)
}
