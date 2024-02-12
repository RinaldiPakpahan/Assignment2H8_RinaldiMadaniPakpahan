package service

import (
	"errors"
	"unit-testing/entity"
	"unit-testing/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	return product, nil
}
func (service ProductService) GetAllProduct() (*[]entity.Product, error) {
	product := service.Repository.FindByAll()

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	return product, nil
}
