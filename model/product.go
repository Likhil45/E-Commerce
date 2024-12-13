package model

type Product struct {
	ProductId    string `json:"product_id" bson:"product_id"`
	ProductName  string `json: "product_name" bson:"name"`
	ProductCost  int64  `json: "product_cost" bson:"cost"`
	ProductColor string `json:"product_color" bson:"color"`
}
