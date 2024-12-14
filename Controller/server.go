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

// Get All
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// GetById
func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range products {
		if item.ProductId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Product{})
}

// Create Product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println("Wrong data sent")
		panic(err)
	}
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// Update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.ProductId == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product model.Product

			err := json.NewDecoder(r.Body).Decode(&product)
			if err != nil {
				panic(err)
			}
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Product{})
}

// Delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	for index, item := range products {
		if item.ProductId == parameter["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
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
	r.HandleFunc("/create", CreateProduct).Methods("POST")
	r.HandleFunc("/{id}", GetProductById)
	r.HandleFunc("/update/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteProduct).Methods("DELETE")
	// r.HandleFunc("/create", CreateProduct).Methods("POST")

	http.ListenAndServe(":8080", r)

}
