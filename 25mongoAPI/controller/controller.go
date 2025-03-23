package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/divy-11/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Admin:jZEOL6CqWnapaRou@cluster0.g8cch4b.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// to connect db
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	//connect db
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Successful.")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection ready!")
}

func addOneMovie(movie model.Netflix) {
	ins, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ins)
}
