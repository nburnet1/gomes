package main

import (
	"fmt"
	// "gomes/api"
	// docs "gomes/docs"

	// "github.com/gin-gonic/gin"

	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           GoMES API
// @version         0.0
// @description     Describes endpoints used by GoMES.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Noah Burnette
// @contact.url    https://github.com/nburnet1/gomes/issues

// @license.name  MIT
// @license.url   https://github.com/nburnet1/gomes/blob/main/LICENSE

// @host      localhost:8080

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	fmt.Println("Hello!!")
	// r := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api/v0"
	// v1 := r.Group("/api/v0")
	// {
	// 	api.RegisterRoutes(v1)
	// }
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	db, err := connectToPostgreSQL()

	if err != nil {
		log.Fatal(err)
	}

	// Perform database migration
	err = db.Table("public.users").AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// r.Run(":8080")

}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=postgres password=password dbname=gomes host=timescaledb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  uint
}
