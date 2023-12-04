package main

import (
	"example.com/mod/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/grandprix", entity.GetGrandPrix)
	router.GET("/grandprix/:id", entity.GetGrandPrixByID)
	router.GET("/categories", entity.GetCategory)
	router.GET("/category/:id", entity.GetCategoryByID)

	router.Run("localhost:8080")
}
