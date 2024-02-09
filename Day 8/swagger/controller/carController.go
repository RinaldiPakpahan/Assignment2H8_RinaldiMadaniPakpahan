package controller

import (
	"fmt"
	"net/http"
	"swagger/database"
	"swagger/models"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @Summary Post details for a given Id
// @Description Post details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars/ [post]
func CreateCar(ctx *gin.Context) {
	var db = database.GetDB()
	var newCar models.Car

	//validate input
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	carInput := models.Car{Merk: newCar.Merk, Harga: newCar.Harga, TypeCars: newCar.TypeCars}
	db.Create(&carInput)

	// newCar.ID = fmt.Sprintf("%d", len(carDatas)+1)
	//models.CarDatas = append(models.CarDatas, newCar)
	ctx.JSON(http.StatusCreated, gin.H{
		"car": carInput,
	})

}

// UpdateCar godoc
// @Summary Update car indetified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars/{id} [PUT]
func UpdateCar(c *gin.Context) {
	var db = database.GetDB()
	var updateCar models.Car

	err := db.First(&updateCar, "Id = ?", c.Param("carID")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found"})
		return
	}

	//validate input
	var valid models.Car
	if err := c.ShouldBindJSON(&valid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&updateCar).Updates(models.Car{Merk: valid.Merk, TypeCars: valid.TypeCars, Harga: valid.Harga})
	fmt.Println(updateCar)

	c.JSON(http.StatusOK, gin.H{
		"message": "car has been successfully updated",
	})
}

// GetData godoc
// @Summary Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetData(ctx *gin.Context) {
	var db = database.GetDB()
	carID := ctx.Param("carID")
	var carData models.Car

	err := db.First(&carData, "Id = ?", carID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error Getting Car Datas ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// DeleteCar godoc
// @Summary Delete car identified by the given Id
// @Description Delete the order corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{Id} [delete]
func DeleteCar(c *gin.Context) {
	var db = database.GetDB()
	//get model if exist
	var carDelete models.Car

	err := db.First(&carDelete, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{
		"message": "car has been successfully deleted",
	})
}
