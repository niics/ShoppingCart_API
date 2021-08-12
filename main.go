package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/products-json", getProductsJSON)
	router.GET("/product/:id", getProductByID)
	router.POST("/products", postProducts)

	router.GET("/orders-json", getOrdersJSON)
	router.GET("/order/:id", getOrderByUID)
	router.POST("/orders/create", createOrder)

	product_init()
	order_init()

	router.Run("localhost:8080")
}
