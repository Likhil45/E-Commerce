package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Likhil45/E-Commerce/model"

	"github.com/Likhil45/E-Commerce/database"
	"github.com/Likhil45/E-Commerce/responses"
	"go.mongodb.org/mongo-driver/bson"
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
