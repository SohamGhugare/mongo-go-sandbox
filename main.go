package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Loading environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env")
	}
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI cannot be empty")
	}

	// Connecting to client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Closing the client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Creating the collection and inserting data
	coll := client.Database("Personal").Collection("students")
	// res, err := coll.InsertOne(context.Background(), bson.M{
	// 	"name": "Soham Ghugare",
	// 	"roll": rand.Intn(9999) + 10000,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// id := res.InsertedID
	// fmt.Printf("Added data at id %v\n", id)

	results := []struct {
		Name string
		Roll int
	}{}
	filter := bson.M{"name": "Soham Ghugare"}
	cur, err := coll.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	defer cur.Close(context.Background())
	if err := cur.All(context.Background(), &results); err != nil {
		panic(err)
	}

	for _, v := range results {
		fmt.Println(v.Roll, ":", v.Name)
	}

}
