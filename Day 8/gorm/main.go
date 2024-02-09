package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	//createUser("rinaldi.pakpahan@gmail.com")

	//getUserById(1)

	//updateData(1, "rinaldi.deff@gmail.com")

	//getUserAll()

	//reateProduct(2, "Iphone", "T")

	//getUserWithProduct()

	deleteProduct(1)

}

func createUser(email string) {
	db := database.GetDB()

	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("New user data: ", user)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return
		}
		print("Error finding user: ", err)
	}

	fmt.Printf("User data: %+v \n", user)
}

func getUserAll() {
	db := database.GetDB()

	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("All Users Data:")
	for _, user := range users {
		fmt.Printf("User Data: %+v\n", user)
	}
}

func updateData(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error createing product data:", err.Error())
		return
	}

	fmt.Println("New product data: ", product)
}

func getUserWithProduct() {
	db := database.GetDB()

	user := models.User{}
	err := db.Preload("Products").Find(&user).Error

	if err != nil {
		fmt.Println("Error getting user data with products: ", err.Error())
		return
	}

	fmt.Println("User data with product")
	fmt.Printf("+v \n", user)

}

func deleteProduct(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product: ", err.Error())
		return
	}

	fmt.Printf("Product with ID %d has been successfully deleted", id)

}
