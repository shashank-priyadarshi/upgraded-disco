package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var db *mongo.Database

func handleRequests() {
	//myRouter := mux.NewRouter().StrictSlash(true)
	//mux.NewRouter().StrictSlash(true)
	//http.HandleFunc("/", homePage)
	http.HandleFunc("/biodata", returnBiodata)
	http.HandleFunc("/todos", returnTodos)
	http.HandleFunc("/articles", returnArticles)
	//myRouter.HandleFunc("/newdata", editData).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnBiodata")
	enableCors(&w)
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("resume")))
}

func returnTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnTodos")
	enableCors(&w)
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("todo")))
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnArticles")
	enableCors(&w)
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("articles")))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Referrer-Policy")
}

func setMongoConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:pass12345@127.0.0.1:27017/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection succeeded:")
	return client
}

func readDataFromCollection(collection *mongo.Collection) []interface{} {
	// Pass these options to the Find method
	findOptions := options.Find()
	//findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []interface{}

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem *interface{}
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, *elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}

func main() {
	client := setMongoConnection()
	db = client.Database("portfolio")

	handleRequests()
}
