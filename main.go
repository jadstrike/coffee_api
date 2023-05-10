package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type coffee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
	ImageURL string `json:"imageurl"`
}

var coffees = []coffee{
	{ID: "1", Name: "FalconX", Category: "Latte", Quantity: 5, ImageURL: "s"},
	{ID: "2", Name: "Rover 23", Category: "Espresso", Quantity: 5, ImageURL: "s"},
	{ID: "3", Name: "Mars", Category: "Mocha", Quantity: 5, ImageURL: "s"},
	{ID: "4", Name: "Venus", Category: "Cappuccino", Quantity: 5, ImageURL: "s"},
}

func getCoffees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, coffees)
}

func coffeeById(c *gin.Context) {
	id := c.Param("id")
	coffee, err := getCoffeeById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "coffee not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, coffee)
}

func getCoffeeById(id string) (*coffee, error) {
	for i, c := range coffees {
		if c.ID == id {
			return &coffees[i], nil
		}
	}
	return nil, errors.New("coffee not found")
}

func createCoffee(c *gin.Context) {
	var newCoffee coffee
	if err := c.BindJSON(&newCoffee); err != nil {
		return
	}
	coffees = append(coffees, newCoffee)
	c.IndentedJSON(http.StatusCreated, newCoffee)
}

func main() {
	router := gin.Default()
	router.GET("/coffees", getCoffees)
	router.GET("/coffee/:id", coffeeById)
	router.POST("/createCoffee", createCoffee)
	router.Run("localhost:8080")

}
