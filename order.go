package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// The ENUM that represent the stage of cart
const (
	PENDING   = iota // 0
	PURCHASED = iota // 1
	CLOSED    = iota // 2
)

// order represents data about a product.
type Order struct {
	PID      string `json:"pid"`
	Quantity int16  `json:"qty"`
}

// order represents data about a product.
type Cart struct {
	UID    string `json:"uid"`
	STAGE  int
	Orders []Order
}

var Carts = []Cart{
	{
		UID:    "0",
		STAGE:  PURCHASED,
		Orders: []Order{},
	},
}

func order_init() {
	fmt.Println("Init order...")
}

// getProducts responds with the list of all products as JSON.
func getOrdersJSON(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Carts)
}

// getOrderByUID locates the album whose UID value matches the user id
// parameter sent by the client, then returns that order as a response.
func getOrderByUID(c *gin.Context) {
	uid := c.Param("uid")

	// Loop through the list of orders, looking for
	// a cart whose user UID value matches the parameter.
	for _, a := range Carts {
		if a.UID == uid {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// createOrder adds an order from JSON received in the request body.
func createOrder(c *gin.Context) {
	var newCart Cart

	// Call BindJSON to bind the received JSON to newProduct.
	if err := c.BindJSON(&newCart); err != nil {
		return
	}

	// Add the new album to the slice.
	Carts = append(Carts, newCart)
	c.IndentedJSON(http.StatusCreated, newCart)
}
