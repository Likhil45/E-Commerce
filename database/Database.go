package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() *mongo.Client {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectToMongoDB()

func FetchProducts(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("E-commerce").Collection(collectionName)

	return collection
}

//Query the collection to get all documents

// cursor, err := collection.Find(ctx, bson.M{})

// if err != nil {
// 	log.Fatal(err)
// }
// defer cursor.Close(ctx)

// //Iterate through collection
// var results []bson.M
// for cursor.Next(ctx) {
// 	var result bson.M
// 	if err := cursor.Decode(&result); err != nil {
// 		log.Fatal(err)
// 	}
// 	results = append(results, result)

// }

// //Check for cursor errors
// if err := cursor.Err(); err != nil {
// 	log.Fatal(err)
// }

// //Print the resutls
// fmt.Println("Products:")
// for _, prod := range results {
// 	fmt.Println(prod)
// }

// defer func() {
// 	if err := client.Disconnect(context.TODO()); err != nil {
// 		panic(err)
// 	}
// }()
