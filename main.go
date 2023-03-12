package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID string `json:id`
	Name string `json:name`
	Author string `json:author`
	Price float64 `json:price`
}

// รายการของหนังสือที่มีทั้งหมด
var books= []book{
	{
		ID: "1", 
		Name: "Harry Potter", 
		Author: "J.K. Rowling", 
		Price: 15.9,
	},
	{
		ID: "2", 
		Name: "One Piece", 
		Author: "Oda Eiichirō", 
		Price: 2.99,
	 },
	{ 
		ID: "3", 
		Name: "demon slayer", 
		Author: "koyoharu gotouge", 
		Price: 2.99,
	},
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	paramID := c.Param("id")
	for _, book := range books {
		if book.ID == paramID {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func AddBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return 
	}
	
	books = append(books, newBook)
	c.JSON(http.StatusOK, newBook)
}



func main() {
	router := gin.Default()
	router.GET("/books", GetBooks)
	router.GET("/book/:id", GetBookByID)
	router.POST("/books", AddBook)
	router.Run("localhost:8080")
}