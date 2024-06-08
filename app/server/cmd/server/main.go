package main

import (
	"fmt"
	"gomes/server/internal/api"
	docs "gomes/server/internal/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"

	"gomes/pkg/model/isa95"
	dao "gomes/server/internal/model/dao/isa95"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
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
	r := setupRouter()

	db, err := connectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	err = initDatabaseSchema(db)
	if err != nil {
		log.Fatal(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./server/internal/model/dao/isa95",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery, // generate mode
	})

	g.UseDB(db) // reuse your gorm db
	g.ApplyBasic(isa95.PersonnelClass{}, isa95.PersonnelClassProp{}, isa95.Person{}, isa95.PersonProp{})

	g.Execute()

	dao.SetDefault(db)

	if dao.Q == nil {
		log.Fatal("query.Q is nil")
	}

	personnelClass, err := dao.Q.PersonnelClass.First()
	if err != nil {
		log.Fatal(err)
	}

	personnelClassProp, err := dao.Q.PersonnelClassProp.First()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(personnelClass.Code)
	fmt.Println(personnelClassProp.Description)
	r.Run(":8080")

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v0"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v0")
	{
		api.RegisterRoutes(v1)
	}

	return r
}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=postgres password=password dbname=gomes host=timescaledb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDatabaseSchema(db *gorm.DB) error {
	return db.AutoMigrate(
		&isa95.PersonnelClassProp{},
		&isa95.PersonnelClassProp{},
		&isa95.Person{},
		&isa95.PersonProp{},
		&isa95.QualificationTest{},
		&isa95.QualificationTestProp{},
	)
}
