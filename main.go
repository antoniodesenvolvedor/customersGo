package main

import (
	"customer-crud/database"
	"github.com/gin-gonic/gin"
	"customer-crud/handlers"
)

func main() {
	session := database.GetSession()
	database.Migrate(session)

	router := gin.Default()
	router.GET("/customers", getAlbums)
	router.Run("localhost:8080")
}