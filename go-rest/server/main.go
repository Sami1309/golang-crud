package main

import (
	"fmt"
	"go-crud-app/middleware"
	"go-crud-app/router"
	"os"
)

func main() {
	//Initialize gin router
	r := router.Router()

	//Establish connection to databse
	middleware.ConnectDB()

	//Run server listening on port 8080

	//os.Setenv("PORT", "8080")

	myport := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(myport)
}
