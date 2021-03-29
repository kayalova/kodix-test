package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kayalova/kodix-test/handler"
	"github.com/kayalova/kodix-test/middleware"
)

func initRoutes(router *gin.Engine) {
	grouped := router.Group("/api/cars")

	grouped.GET("/", middleware.TransformGetInput, handler.GetAll)
	grouped.POST("/create", middleware.IsValidCreateInput, handler.CreateOne)
	grouped.PUT("/update", middleware.IsValidDeleteUpdateInput, handler.UpdateOne)
	grouped.DELETE("/delete", middleware.IsValidDeleteUpdateInput, handler.DeleteOne)

}

// Run server
func Run(port string) {
	router := gin.Default()
	initRoutes(router)

	router.Run(":" + port)
}
