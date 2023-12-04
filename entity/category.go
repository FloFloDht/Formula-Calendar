package entity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// init the Category struct
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var category = []Category{
	{
		ID:   "1",
		Name: "Formula 1",
	},
	{
		ID:   "2",
		Name: "Formula 2",
	},
	{
		ID:   "3",
		Name: "Formula 3",
	},
}

// GetCategory responds with the list of all Category JSON
func GetCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, category)
}

// GetCategoryByID locates the Category whose ID value matches the id
// parameter sent by the client, then returns that Category as a response.
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of Category, looking for
	// a Category whose ID value matches the parameter.
	for _, a := range category {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "category not found"})
}
