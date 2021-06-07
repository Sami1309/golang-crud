package middleware

import (
	//remote packages
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//my packages
	"go-crud-app/models"
)

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publisher   string `json:"publisher" binding:"required"`
	PublishDate string `json:"publish date" binding:"required"` //Format: YYYY/DD/MM
	Rating      uint8  `json:"rating" binding:"required"`       //1, 2, or 3
	CheckedOut  bool   `json:"checked out"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish date"`
	Rating      uint8  `json:"rating"`
	CheckedOut  bool   `json:"checked out"`
}

func ClearTable() {
	// DB.Delete(&models.Book{})
	DB.Exec("DELETE FROM Books; UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='Books';")
}

func CreateBook(c *gin.Context) {
	// Validate input
	var newBook CreateBookInput

	fmt.Println(c)
	invalid_request := c.ShouldBindJSON(&newBook)

	if invalid_request != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": invalid_request.Error()})
		return
	}

	//validate book input here to abide by formatting, or perhaps way to do it in database?
	valid_book := len(strings.Split(newBook.PublishDate, "-")) == 3 &&
		(newBook.Rating >= 1 && newBook.Rating <= 3)

	if valid_book == false {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input."})
		return
	}

	// Create book
	book := models.Book{Title: newBook.Title, Author: newBook.Author, Publisher: newBook.Publisher, PublishDate: newBook.PublishDate, Rating: newBook.Rating, CheckedOut: newBook.CheckedOut}
	DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"Created book": book})
}

// GET /books
// Find all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"Books": books})
}

func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	id := c.Param("id")
	//400 bad request if not an integer?

	query := DB.Where("id = ?", id).First(&book)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Book does not exist."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Found Book": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book

	id := c.Param("id")
	//400 bad request if not an integer?

	query := DB.Where("id = ?", id).First(&book)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Book not found."})
		return
	}

	// Validate input
	var updateBook UpdateBookInput

	invalid_request := c.ShouldBindJSON(&updateBook)

	//checkbookinput(newBook) //Ensures input follows proper formatting/data types

	if invalid_request != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": invalid_request.Error()})
		return
	}

	DB.Model(&book).Updates(updateBook)

	c.JSON(http.StatusOK, gin.H{"Updated book": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book

	id := c.Param("id")
	query := DB.Where("id = ?", id).First(&book)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Book not found."})
		return
	}

	DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"Deleted Book": true})
}

func DeleteBooks(c *gin.Context) {
	ClearTable()
	c.JSON(http.StatusOK, gin.H{"Deleted Books": true})
}
