package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserName     string             `json:"userName" bson:"user_name"`
	UserId       primitive.ObjectID `json:"userId" bson:"_id"`
	UserPassword string             `json:"userPassword" bson:"user_password"`
	UserEmail    string             `json:"userEmail" bson:"user_email"`
	UserAddress  string             `json:"userAddress" bson:"user_address"`
}
