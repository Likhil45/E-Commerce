package main

import (
	routes "github.com/Likhil45/E-Commerce/Routes"
	"github.com/Likhil45/E-Commerce/database"
)

func main() {
	database.ConnectToMongoDB()
	routes.StartServer()

}
