package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"

	"net/http"
	"time"

	"github.com/Likhil45/E-Commerce/database"
	model "github.com/Likhil45/E-Commerce/model"
	"github.com/Likhil45/E-Commerce/responses"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var products []model.Product
var productCollection *mongo.Collection = database.Fetch(database.DB, "products")

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
}

// GetById
func GetProductById(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var product model.Product
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(params["id"])

	err := productCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&product)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
	}
	rw.WriteHeader(http.StatusOK)
	response := responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": product}}
	json.NewEncoder(rw).Encode(response)
}

// Create Product
func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var product model.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var validate = validator.New()
	defer cancel()

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println("Wrong data sent")
		panic(err)
	}
	if validationErr := validate.Struct(&product); validationErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	newproduct := model.Product{
		ProductId:    primitive.NewObjectID(),
		ProductName:  product.ProductName,
		ProductCost:  product.ProductCost,
		ProductColor: product.ProductColor,
	}
	result, err := productCollection.InsertOne(ctx, newproduct)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	response := responses.ProductResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(rw).Encode(response)

}

// Update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	prodId := params["id"]
	ObjId, _ := primitive.ObjectIDFromHex(prodId)
	for index, item := range products {
		if item.ProductId == ObjId {
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
	prodId := parameter["id"]
	ObjId, _ := primitive.ObjectIDFromHex(prodId)
	for index, item := range products {
		if item.ProductId == ObjId {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
