package main

import (
	controller "github.com/Likhil45/E-Commerce/Controller"
	"github.com/Likhil45/E-Commerce/database"
)

func main() {
	database.ConnectToMongoDB()
	controller.StartServer()

}
