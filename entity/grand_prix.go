package entity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// init the GrandPrix struct
type GrandPrix struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Date     string `json:"date"`
	Name     string `json:"name"`
}

var grand_prix = []GrandPrix{
	{
		ID:       "1",
		Location: "Sakir",
		Date:     "2024-03-02",
		Name:     "Bahrain Grand Prix",
	},
	{
		ID:       "2",
		Location: "Djeddah",
		Date:     "2024-03-19",
		Name:     "Saudi Arabian Grand Prix",
	},
	{
		ID:       "3",
		Location: "Melbourne",
		Date:     "2024-03-24",
		Name:     "Australian Grand Prix",
	},
}

// getGrandPrix responds with the list of all GrandPrix JSON
func GetGrandPrix(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, grand_prix)
}

// getGrandPrixByID locates the GrandPrix whose ID value matches the id
// parameter sent by the client, then returns that GrandPrix as a response.
func GetGrandPrixByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of GrandPrix, looking for
	// a GrandPrix whose ID value matches the parameter.
	for _, a := range grand_prix {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "grand prix not found"})
}
