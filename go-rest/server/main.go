package main

import (
	"go-crud-app/middleware"
	"go-crud-app/router"
)

func main() {
	//Initialize gin router
	r := router.Router()

	//Establish connection to databse
	middleware.ConnectDB()

	//Run server listening on port 8080
	r.Run(":8080")
}
