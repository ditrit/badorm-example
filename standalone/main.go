package main

import (
	"time"

	"github.com/ditrit/badaas/badorm"
	"github.com/ditrit/badorm-example/standalone/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := NewGormDBConnection()
	if err != nil {
		panic(err)
	}

	err = gormDB.AutoMigrate(
		models.Product{},
		models.Company{},
		models.Seller{},
		models.Sale{},
	)
	if err != nil {
		panic(err)
	}

	crudProductService, crudProductRepository := badorm.GetCRUD[models.Product, badorm.UUID](gormDB)

	CreateCRUDObjects(gormDB, crudProductRepository)
	QueryCRUDObjects(crudProductService)
	// TODO ejemplo unsafe
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
