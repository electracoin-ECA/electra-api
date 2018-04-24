package main

import (
	"github.com/Electra-project/electra-api/src/api"
	"github.com/Electra-project/electra-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.CORS())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	v1 := router.Group("/v1")
	{
		v1.GET("/price/:currency", api.GetPrice)
	}

	router.Run()
}
