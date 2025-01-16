package routes

import (
	"net/http"

	"github.com/Likhil45/E-Commerce/controller"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	//Define Routes

	//Products
	r.HandleFunc("/", controller.HomeHandler).Methods("GET")
	r.HandleFunc("/test", controller.TestHandler)
	r.HandleFunc("/all", controller.GetAllProducts).Methods("GET")
	r.HandleFunc("/create", controller.CreateProduct).Methods("POST")
	r.HandleFunc("/{id}", controller.GetProductById)
	r.HandleFunc("/update/{id}", controller.UpdateProduct).Methods("PUT")
	r.HandleFunc("/delete/{id}", controller.DeleteProduct).Methods("DELETE")

	//Users
	r.HandleFunc("/user/all", controller.GetAllUser).Methods("GET")
	r.HandleFunc("/user/create", controller.CreateUser).Methods("GET")

	http.ListenAndServe(":8080", r)

}
