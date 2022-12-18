package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

var db *mongo.Database

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/biodata", returnBiodata).Methods("GET")
	r.HandleFunc("/todos", returnTodos).Methods("GET")
	r.HandleFunc("/articles", returnArticles).Methods("GET")
	return r
}

//func originValidator(origin string) bool {
//	valid := false
//	err := pool.QueryRow("SELECT IF(origin=?, True, False) as 'valid' FROM origins", origin).Scan(&valid)
//	if err != nil {
//		return false
//	}
//	return valid
//}

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	//ttl := handlers.MaxAge(3600)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("API_PORT")), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnBiodata")
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("resume")))
}

func returnTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnTodos")
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("todo")))
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnArticles")
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("articles")))
}

func setMongoConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MongoURI"))

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
	db = client.Database(os.Getenv("DB_NAME"))

	handleRequests()
}
