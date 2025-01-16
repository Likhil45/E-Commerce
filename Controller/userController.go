package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Likhil45/E-Commerce/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/Likhil45/E-Commerce/database"
	"github.com/Likhil45/E-Commerce/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.Fetch(database.DB, "user")

// GetALL
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []model.User
	defer cancel()

	//Error Handling
	res, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var user model.User
		if err = res.Decode(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Unable to Decode ", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	response := responses.ProductResponse{Status: http.StatusOK, Message: "Successfully Fetched all Users", Data: map[string]interface{}{"users": users}}
	json.NewEncoder(w).Encode(response)

}

// Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	var validate = validator.New()
	defer cancel()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Wrong data is sent")
		panic(err)
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	newUser := model.User{
		UserName:     user.UserName,
		UserId:       primitive.NewObjectID(),
		UserPassword: user.UserPassword,
		UserEmail:    user.UserEmail,
		UserAddress:  user.UserAddress,
	}
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := responses.ProductResponse{Status: http.StatusCreated, Message: "Succefully Added", Data: map[string]interface{}{"Added": result}}
	json.NewEncoder(w).Encode(response)

}

// Get user by id
func GetUserById(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(params["id"])

	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
	}
	rw.WriteHeader(http.StatusOK)
	response := responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}}
	json.NewEncoder(rw).Encode(response)
}
