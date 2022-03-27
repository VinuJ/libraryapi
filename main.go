package main

import (
	// "errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct{
	ID		string  `json:"id"`
	Title	string  `json:"title"`
	Author	string  `json:"author"`
	Quantity  int   `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Skulduggery Pleasant", Author: "Derek Landy", Quantity: 4},
	{ID: "2", Title: "Harry Potter", Author: "J. K. Rowling", Quantity: 5},
	{ID: "3", Title: "Goosebumps", Author: "R. L. Stine", Quantity: 9},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}