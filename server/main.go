package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type schedule struct {
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}

var db *mongo.Database

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/biodata", returnBiodata).Methods("GET")
	r.HandleFunc("/todos", returnTodos).Methods("GET")
	r.HandleFunc("/articles", returnArticles).Methods("GET")
	r.HandleFunc("/githubdata", returnGitHubData).Methods("GET")
	r.HandleFunc("/schedule", writeNewSchedule).Methods("POST")
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
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
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

func returnGitHubData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnGitHubData")
	json.NewEncoder(w).Encode(readDataFromCollection(db.Collection("githubdata")))
}

func writeNewSchedule(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: writing articles")
	reqStatus := writeDataToCollection(db.Collection("schedule"), req)
	json.NewEncoder(w).Encode(reqStatus)
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
	var sortDoc interface{} = make(map[string]interface{})
	sortDoc.(map[string]interface{})["_id"] = -1
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetSort(sortDoc)
	findOptions.SetLimit(1)

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

func writeDataToCollection(collection *mongo.Collection, req *http.Request) *http.ResponseWriter {
	var response *http.ResponseWriter
	fmt.Println("collection: ", collection)
	fmt.Println("request: ")
	var ifa schedule
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error while reading request body: ", err)
	}
	err = json.Unmarshal(reqBody, &ifa)
	if err != nil {
		fmt.Println("Error while unmarshaling request body: ", err)
	} else {
		fmt.Println(ifa)
		insertResult, err := collection.InsertOne(context.TODO(), ifa)
		if err != nil {
			fmt.Println("Error while Inserting new schedule to MongoDB")
		} else {
			fmt.Println("New schedule saved successfully! ", insertResult)
		}
	}
	return response
}

func StartServer() {
	client := setMongoConnection()
	db = client.Database(os.Getenv("DB_NAME"))

	handleRequests()
}
