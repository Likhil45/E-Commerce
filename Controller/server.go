package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"github.com/Likhil45/E-Commerce/database"
	model "github.com/Likhil45/E-Commerce/model"
	"github.com/Likhil45/E-Commerce/responses"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var products []model.Product
var productCollection *mongo.Collection = database.FetchProducts(database.DB, "products")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Test Page")
}

// Get All
func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var products []model.Product
	defer cancel()

	//Error Handling
	results, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}
	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleProduct model.Product
		if err = results.Decode(&singleProduct); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
		}

		products = append(products, singleProduct)

	}
	rw.WriteHeader(http.StatusOK)
	response := responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}}
	json.NewEncoder(rw).Encode(response)
	// prod := database.FetchProducts(productCollection,"products")
	// json.NewEncoder(w).Encode(&productCollection)
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
	r.HandleFunc("/all", GetAllProducts).Methods("GET")
	r.HandleFunc("/create", CreateProduct).Methods("POST")
	r.HandleFunc("/{id}", GetProductById)
	r.HandleFunc("/update/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteProduct).Methods("DELETE")
	// r.HandleFunc("/create", CreateProduct).Methods("POST")

	http.ListenAndServe(":8080", r)

}
