package main

import (
	"log"
	"time"

	"github.com/ditrit/badaas/badorm"
	"github.com/ditrit/badorm-example/models"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		fx.Provide(NewLogger),
		// DB modules
		fx.Provide(NewGORMDBConnection),
		// activate BaDORM
		badorm.BaDORMModule,

		// start example data
		badorm.GetCRUDServiceModule[models.Company, uuid.UUID](),
		badorm.GetCRUDServiceModule[models.Product, uuid.UUID](),
		badorm.GetCRUDServiceModule[models.Seller, uuid.UUID](),
		badorm.GetCRUDServiceModule[models.Sale, uuid.UUID](),

		fx.Provide(CreateCRUDObjects),
		fx.Invoke(QueryCRUDObjects),
	).Run()
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func NewGORMDBConnection() (*gorm.DB, error) {
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

		log.Println("Database connection failed with error %q", err.Error())
		log.Println(
			"Retrying database connection %d/%d in %ds",
			numberRetry+1, retryAmount, retryTime,
		)
		time.Sleep(time.Duration(retryTime) * time.Second)
	}

	return nil, err
}
