package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

var pokemon []Pokemon

func get(c *gin.Context) {
	c.JSON(http.StatusOK, pokemon)
}

func create(c *gin.Context) {
	var rs Pokemon
	if err := c.ShouldBindJSON(&rs); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rs.ID = strconv.Itoa(len(pokemon) + 1)
	pokemon = append(pokemon, rs)
	c.JSON(http.StatusCreated, rs)
}

func update(c *gin.Context) {
	id := c.Param("id")
	var rs Pokemon
	if err := c.ShouldBindJSON(&rs); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for i, pkm := range pokemon {
		if pkm.ID == id {

			rs.ID = pkm.ID
			pokemon[i] = rs
			c.JSON(http.StatusOK, gin.H{"message": "Update pokemon successfully"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"error": "Pokemon not found"})
}

func delete(c *gin.Context) {
	id := c.Param("id")

	for i, pk := range pokemon {
		if pk.ID == id {
			pokemon = append(pokemon[:i], pokemon[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted pokemon successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
}

func main() {
	r := gin.Default()

	// Get method
	r.GET("/", get)

	// Get method
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	// Post method
	r.POST("/", create)

	// Put method
	r.PUT("/:id", update)

	// Delete method
	r.DELETE("/:id", delete)

	r.Run(":8080")
}
