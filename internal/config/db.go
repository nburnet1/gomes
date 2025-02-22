package config

import (
	"github.com/nburnet1/gomes/pkg/namespace"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("./gomes.db"), &gorm.Config{})

func MigrateFromRegistry() error{
	var err error
	modelRegistry := namespace.GetModelRegistry()
	for _, model := range modelRegistry {
		err = DB.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}


