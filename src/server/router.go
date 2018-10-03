package server

import (
	"github.com/Electra-project/electra-api/src/controllers"
	"github.com/Electra-project/electra-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

// Router binds the routes to the controllers.
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Electra Auth API",
			"version": "1.0",
		})
	})

	priceController := new(controllers.PriceController)
	router.GET("/v1/priceother/:coin/:currency", priceController.Get) // Legacy route
	router.GET("/v1/price/:currency", priceController.Get)            // Legacy route
	router.GET("/price/:currency", priceController.Get)
	router.GET("/priceother/:coin/:currency", priceController.Get)
	statusController := new(controllers.StatusController)
	router.GET("/status", statusController.Get)

	addrController := new(controllers.AddressController)
	router.GET("/address/:hash", addrController.Get)
	router.GET("/address/:hash/transactions", addrController.GetTransactions)
	transactionController := new(controllers.TransactionController)
	router.GET("/transaction/:id", transactionController.Get)
	router.POST("/transaction", transactionController.Post)

	userGroup := router.Group("user")
	{
		userTokenController := new(controllers.UserTokenController)
		userGroup.GET("/:purseHash/token", userTokenController.Get)
		userGroup.POST("/:purseHash/token", userTokenController.Post)
	}

	router.Use(middlewares.IsUser())
	{
		userController := new(controllers.UserController)
		router.GET("/user", userController.Get)
		router.POST("/user", userController.Post)
		router.PUT("/user", userController.Put)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
