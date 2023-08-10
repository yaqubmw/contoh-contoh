package repository_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"your-app/model"
	"your-app/repository"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repository.ProductRepository
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	// Connect to a PostgreSQL test database
	dsn := "user=username password=password dbname=testdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		suite.T().Fatalf("error opening database: %s", err)
	}
	suite.db = db

	// Migrate the database schema
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		suite.T().Fatalf("error migrating database schema: %s", err)
	}

	suite.repo = repository.NewProductRepository(db)
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	// Drop the test database table
	err := suite.db.Migrator().DropTable(&model.Product{})
	if err != nil {
		suite.T().Errorf("error dropping table: %s", err)
	}
	suite.db.Close()
}

func (suite *ProductRepositoryTestSuite) TestCreateNewProduct_Success() {
	// Create a new product
	product := model.Product{
		Name:  "Product A",
		Price: 10000,
		Uom:   model.Uom{ID: "1"},
	}

	err := suite.repo.Create(&product)
	assert.NoError(suite.T(), err)

	// Fetch the product from the database and verify its values
	var fetchedProduct model.Product
	err = suite.db.First(&fetchedProduct, "id = ?", product.ID).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), product.Name, fetchedProduct.Name)
	assert.Equal(suite.T(), product.Price, fetchedProduct.Price)
}

func (suite *ProductRepositoryTestSuite) TestListProducts_Success() {
	// Insert dummy products into the database for testing
	dummyProducts := []model.Product{
		{Name: "Product A", Price: 10000, Uom: model.Uom{ID: "1"}},
		{Name: "Product B", Price: 20000, Uom: model.Uom{ID: "2"}},
	}

	for i := range dummyProducts {
		err := suite.db.Create(&dummyProducts[i]).Error
		assert.NoError(suite.T(), err)
	}

	// Retrieve the list of products
	products, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), products, len(dummyProducts))
}

// Add more test cases as needed...

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
