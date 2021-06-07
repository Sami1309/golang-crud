package router

import (
	"github.com/gin-gonic/gin"

	"go-crud-app/middleware"
)

func Router() *gin.Engine {
	//Initialize gin router
	router := gin.Default()

	//Add book entry to database
	router.POST("/books", middleware.CreateBook)
	//Retrieve book entries from database
	router.GET("/books", middleware.FindBooks)
	router.GET("/books/:id", middleware.FindBook)
	//Update book entry already in database
	router.PATCH("/books/:id", middleware.UpdateBook)
	//Remove book entry from database
	router.DELETE("/books/:id", middleware.DeleteBook)
	//Remove all books from database
	router.DELETE("/books", middleware.DeleteBooks)

	return router

}
