package router

import (
	"swagger/controller"
	_ "swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description this is a sample service for managing carDatas
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controller.CreateCar)

	router.PUT("/cars/:carID", controller.UpdateCar)

	router.GET("/cars/:carID", controller.GetData)

	router.DELETE("/cars/:id", controller.DeleteCar)

	router.GET("/cars", controller.GetAllCars)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
