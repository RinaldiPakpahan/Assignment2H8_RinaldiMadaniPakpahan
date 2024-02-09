package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key`
	Email     string `gorm:"not null;unique;type:varchar(191)`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint   `gorm:"primary_key`
	Name      string `gorm:"not null;unique;type:varchar(191)"`
	Brand     string `gorm:"not null;unique;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreateProduct(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}

	return
}
