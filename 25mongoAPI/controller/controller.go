package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/divy-11/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
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

// functions to be used in controllers
func addOneMovie(movie model.Netflix) {
	ins, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ins)
}

func setMovieWatched(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	updateInfo := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, updateInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count:", res.ModifiedCount)
}

func delOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.DeletedCount)
}

func deleteAll() int64 {
	res, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return res.DeletedCount
}

func allMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{})
	defer cur.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// CONTROLLERS:
func AllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	respMovies := allMovies()
	json.NewEncoder(w).Encode(respMovies)
}

func CreatingMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	addOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func WatchedMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)

	setMovieWatched(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	delOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	cnt := deleteAll()
	json.NewEncoder(w).Encode(cnt)
}
