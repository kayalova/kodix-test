package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsValidCreateInput /api/cars/create body
func IsValidCreateInput(c *gin.Context) {
	if c.PostForm("Brand") == "" || c.PostForm("Model") == "" || c.PostForm("Price") == "" || c.PostForm("Status") == "" || c.PostForm("Milage") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't create a car. Please provide valid body"})
		c.Abort()
	}
}

// TransformGetInput transforms query string: extract value from query key value (which is array)
func TransformGetInput(c *gin.Context) {
	params := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		params[k] = v[0]
	}

	c.Set("searchParams", params)
}

// IsValidDeleteUpdateInput /api/cars/delete query and /api/cars/update query
func IsValidDeleteUpdateInput(c *gin.Context) {
	id, isExists := c.GetQuery("id")
	// add db request
	if id == "" || !isExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid id"})
		c.Abort()
	}

	// car := repository.FindById()
	// if !car  {
	// c.JSON(http.StatusBadRequest, gin.H{"message": "Can't find a car. Please provide valid id"})
	// c.Abort(
	// }
}

func isValidDeleteInput(c *gin.Context) {
	id, isExists := c.GetQuery("id")
	// add db request
	if id == "" || !isExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid id"})
		c.Abort()
	}

	// car := repository.FindById(id)
	// if !car {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Can not find car with provided id"})
	// 	c.Abort()
	// }

}
