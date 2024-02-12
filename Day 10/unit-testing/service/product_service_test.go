package service

import (
	"testing"
	"unit-testing/entity"
	"unit-testing/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {

	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product Not Found", err.Error(), "Error has to be `Product Not Fount`")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kacamata",
	}
	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.NotNil(t, result)
	assert.Nil(t, err)

	assert.Equal(t, product.Id, result.Id, "Result has to be `2`")
	assert.Equal(t, product.Name, result.Name, "Result has to be `kacamata`")
	assert.Equal(t, &product, result, "Product has to be a product data with Id `2`")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	product := []entity.Product{
		{
			Id:   "2",
			Name: "Kacamata",
		},
		{
			Id:   "1",
			Name: "Buku",
		},
	}
	productRepository.Mock.On("FindByAll").Return(product)

	result, err := productService.GetAllProduct()

	for i, v := range *result {
		assert.Equal(t, product[i].Name, v.Name)
	}

	assert.NotNil(t, result)
	assert.Nil(t, err)

	assert.Equal(t, &product, result)
}
func TestProductServiceGetAllProductFail(t *testing.T) {
	product := []entity.Product{
		{
			Id:   "2",
			Name: "Kacamata",
		},
		{
			Id:   "1",
			Name: "Buku",
		},
	}
	productRepository.Mock.On("FindByAll").Return(product)

	result, _ := productService.GetAllProduct()

	assert.Equal(t, len(product), len(*result)+1)
}
