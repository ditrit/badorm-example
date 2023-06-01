package main

import (
	"time"

	"github.com/ditrit/badaas/badorm"
	"github.com/ditrit/badorm-example/fx/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		// DB modules
		fx.Provide(NewGormDBConnection),
		// activate BaDORM
		fx.Provide(GetModels),
		badorm.BaDORMModule,

		// start example data
		badorm.GetCRUDServiceModule[models.Company](),
		badorm.GetCRUDServiceModule[models.Product](),
		badorm.GetCRUDServiceModule[models.Seller](),
		badorm.GetCRUDServiceModule[models.Sale](),

		fx.Provide(CreateCRUDObjects),
		fx.Invoke(QueryCRUDObjects),
	).Run()
}

func NewGormDBConnection() (*gorm.DB, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return badorm.ConnectToDialector(
		logger,
		badorm.CreatePostgreSQLDialector("localhost", "root", "postgres", "disable", "badaas_db", 26257),
		10, time.Duration(5)*time.Second,
	)
}

// TODO el depency injector de google

func GetModels() badorm.GetModelsResult {
	return badorm.GetModelsResult{
		Models: []any{
			models.Product{},
			models.Company{},
			models.Seller{},
			models.Sale{},
		},
	}
}
