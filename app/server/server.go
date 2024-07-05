package server

import (
	"gomes/pkg/err"
	"gomes/server/internal/api"
	docs "gomes/server/internal/api/docs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	gomesGen "gomes/pkg/gen"
	ref "gomes/server/internal/model/gen"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

//	@title			GoMES API
//	@version		0.0
//	@description	Describes endpoints used by GoMES.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Noah Burnette
//	@contact.url	https://github.com/nburnet1/gomes/issues

//	@license.name	GPL v3.0
//	@license.url	https://github.com/nburnet1/gomes/blob/main/LICENSE

//	@host	localhost:8080

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func Start(serverError *err.Error) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	modelGen, err := gomesGen.NewModelGenerationFromFile("./server/config/modelGeneration.json")
	serverError.SetErrorAndFatalIfNotNil(err)

	err = modelGen.SetModelsToGenerateAndWriteToFile()
	serverError.SetErrorAndFatalIfNotNil(err)

	db, err := connectToPostgreSQL()
	serverError.SetErrorAndFatalIfNotNil(err)

	daoGenerator := gen.NewGenerator(gen.Config{
		OutPath: "./server/internal/dao/isa95/",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})

	apiEngine := setupRouter()

	server := serverInstance{
		db:           db,
		logger:       logger,
		daoGenerator: daoGenerator,
		apiEngine:    apiEngine,
	}

	err = server.initDatabaseSchema()
	serverError.SetErrorAndFatalIfNotNil(err)

	server.daoGenerator.UseDB(db)
	for _, ref := range ref.GeneratedModelReferences {
		server.daoGenerator.ApplyBasic(ref)
	}
	server.daoGenerator.Execute()

	server.apiEngine.Run(":8080")

}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=postgres password=password dbname=gomes host=timescaledb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v0"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v0 := r.Group("/api/v0")
	{
		api.RegisterRoutes(v0)
	}

	return r
}



type serverInstance struct {
	db           *gorm.DB
	logger       *zap.Logger
	daoGenerator *gen.Generator
	apiEngine    *gin.Engine
}

func (si *serverInstance) initDatabaseSchema() error {
	for _, ref := range ref.GeneratedModelReferences {
		err := si.db.AutoMigrate(&ref)
		if err != nil {
			return err
		}
	}
	return nil

}


