package main

import (
	"fmt"
	"log"

	"github.com/ditrit/badaas/badorm"
	"github.com/ditrit/badorm-example/models"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func CreateCRUDObjects(
	db *gorm.DB,
	crudProductRepository badorm.CRUDRepository[models.Product, uuid.UUID],
) []*models.Product {
	log.Println("Setting up CRUD example")

	product1 := &models.Product{
		Int: 1,
	}
	_ = crudProductRepository.Create(db, product1)

	product2 := &models.Product{
		Int: 2,
	}
	_ = crudProductRepository.Create(db, product2)

	company1 := &models.Company{
		Name: "ditrit",
	}
	_ = db.Create(company1).Error
	company2 := &models.Company{
		Name: "orness",
	}
	_ = db.Create(company2).Error

	seller1 := &models.Seller{
		Name:      "franco",
		CompanyID: &company1.ID,
	}
	_ = db.Create(seller1).Error
	seller2 := &models.Seller{
		Name:      "agustin",
		CompanyID: &company2.ID,
	}
	_ = db.Create(seller2).Error

	sale1 := &models.Sale{
		Product: product1,
		Seller:  seller1,
	}
	_ = db.Create(sale1).Error
	sale2 := &models.Sale{
		Product: product2,
		Seller:  seller2,
	}
	_ = db.Create(sale2).Error

	log.Println("Finished creating CRUD example")

	return []*models.Product{product1, product2}
}

func QueryCRUDObjects(
	_ []*models.Product,
	crudProductService badorm.CRUDService[models.Product, uuid.UUID],
	shutdowner fx.Shutdowner,
) {
	log.Println("Products with int = 1 are:")
	result, err := crudProductService.GetEntities(
		map[string]any{
			"int": 1.0,
		},
	)
	if err != nil {
		log.Panicln(err)
	}
	for _, product := range result {
		fmt.Printf("%+v\n", product)
	}
	shutdowner.Shutdown()
}
