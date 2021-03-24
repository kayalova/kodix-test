package controller

import "github.com/gin-gonic/gin"

func initRoutes(router *gin.Engine) {
	// grouped := router.Group("/api/cars")

	// grouped.GET("/", handlers.GetAll)
	// grouped.POST("/create", handlers.CreateOne)
	// grouped.PUT("/update", handlers.UpdateOne)
	// grouped.DELETE("/update", handlers.UpdateOne)

}

// Run server
func Run(port string) {
	router := gin.Default()
	initRoutes(router)

	router.Run(":" + port)
}
