package auth

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// Initialize the database connection
func InitDB() {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		// Auto Migrate the User model
		DB.AutoMigrate(&User{})
	})
}