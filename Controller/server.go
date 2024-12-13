package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/Likhil45/E-Commerce/Model"
	"github.com/gorilla/mux"
)

var products []model.Product

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Test Page")
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}
func StartServer() {
	r := mux.NewRouter()

	products = append(products, model.Product{ProductId: "1", ProductName: "Chair", ProductColor: "Brown", ProductCost: 50})
	products = append(products, model.Product{ProductId: "2", ProductName: "Sofa", ProductColor: "Grey", ProductCost: 100})

	//Define Routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/test", TestHandler)
	r.HandleFunc("/all", GetAllProducts)

	http.ListenAndServe(":8080", r)

}
