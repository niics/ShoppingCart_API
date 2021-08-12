package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// product represents data about a product.
type product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"des"`
	Price       float64 `json:"price"`
	Quantity    int16   `json:"qty"`
}

// products slice to seed record product data.
var products = []product{
	{ID: "0", Name: "Test Goods", Description: "The first good for test.", Price: 9.99, Quantity: 10},
}

func product_init() {
	fmt.Println("Init product...")
}

// getProducts responds with the list of all products as JSON.
func getProductsJSON(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// getProductByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that product as a response.
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of products, looking for
	// an album whose ID value matches the parameter.
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// postProducts adds an product from JSON received in the request body.
func postProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}
