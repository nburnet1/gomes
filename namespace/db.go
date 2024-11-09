package namespace

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("gomes.db"), &gorm.Config{})

func MigrateFromRegistry() error{
	var err error
	for _, model := range modelRegistry {
		err = DB.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}


