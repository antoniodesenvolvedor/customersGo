package main

import (
	"customer-crud/configs"
	"customer-crud/database"
	_ "customer-crud/docs"
	"customer-crud/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Este é o título da minha api
// @version 1.0
// @description Descrição da minha api

// @BasePath /v1
// @query.collection.format multi

// @x-extension-openapi {"example": "value on a json format"}


func main() {
	configs.LoadEnvFromFile()

	session := database.GetSession()
	database.Migrate(session)

	router := gin.Default()
	{
		v1 := router.Group("/v1")
		{
			customers := v1.Group("/customers")
			customers.POST("/", handlers.CreateCustomer)
			customers.GET("/:id", handlers.GetCustomer)
			customers.GET("/", handlers.GetCustomers)
			customers.PUT("/:id", handlers.UpdateCustomer)
			customers.DELETE("/:id", handlers.DeleteCustomer)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(fmt.Sprintf("%s:%s", configs.GetConfig().AppDefaultHost,  configs.GetConfig().AppDefaultPort))
}