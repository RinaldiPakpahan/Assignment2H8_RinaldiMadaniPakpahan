package repository

import "unit-testing/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindByAll() *[]entity.Product
}
