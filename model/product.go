package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ProductId    primitive.ObjectID `json:"product_id,omitempty" bson:"_id"`
	ProductName  string             `json:"product_name,omitempty" validate:"required" bson:"name"`
	ProductCost  int64              `json:"product_cost,omitempty" validate:"required" bson:"cost"`
	ProductColor string             `json:"product_color,omitempty" validate:"required" bson:"color"`
}
