package main

import (
	"log"
	"time"

	"github.com/ditrit/badaas/badorm"
	"github.com/ditrit/badorm-example/fx/models"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
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
	dsn := "user=root password=postgres host=localhost port=26257 sslmode=disable dbname=badaas_db"
	var err error
	retryAmount := 10
	retryTime := 5
	for numberRetry := 0; numberRetry < retryAmount; numberRetry++ {
		database, err := gorm.Open(postgres.Open(dsn))
		if err == nil {
			log.Println("Database connection is active")
			return database, nil
		}

		log.Printf("Database connection failed with error %q\n]", err.Error())
		log.Printf(
			"Retrying database connection %d/%d in %ds\n",
			numberRetry+1, retryAmount, retryTime,
		)
		time.Sleep(time.Duration(retryTime) * time.Second)
	}

	return nil, err
}

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
