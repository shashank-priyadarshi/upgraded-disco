package mongoconnection

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"server/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	done = make(chan bool)
	db   *mongo.Database
)

func setMongoConnection() {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.FetchConfig().MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("MongoDB connection succeeded:")

	go func() {
		// Wait for the task to finish or for a timeout
		select {
		case <-done:
			// Task finished, close the MongoDB connection
			err := client.Disconnect(context.TODO())
			if err != nil {
				log.Println("Error while closing connection to mongo after task completion: ", err)
			}
			fmt.Println("MongoDB connection closed after task completion")
		case <-time.After(20 * time.Second):
			// Timeout, close the MongoDB connection
			err := client.Disconnect(context.TODO())
			if err != nil {
				log.Println("Error while closing connection to mongo due to timeout: ", err)
			}
			fmt.Println("MongoDB connection closed due to timeout")
		}
	}()

	db = client.Database(config.FetchConfig().DBNAME)
}

func ReadDataFromCollection(collection string) []byte {
	setMongoConnection()
	var sortDoc interface{} = make(map[string]interface{})
	sortDoc.(map[string]interface{})["_id"] = -1
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetSort(sortDoc)
	findOptions.SetLimit(1)

	// Here's an array in which you can store the decoded documents
	var results []byte
	cur, err := db.Collection(collection).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem *interface{}
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}

		if b, ok := (*elem).(primitive.D); ok {
			marshalledJson, err := json.Marshal(b.Map())
			if err != nil {
				log.Println("error while marshaling mongo data: ", err)
			}
			results = append(results, marshalledJson...)
		}

		// results = append(results, *elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	done <- true
	return results
}

func WriteDataToCollection(collectionName string, data interface{}) error {
	setMongoConnection()
	collection := db.Collection(collectionName)

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("Error while Inserting github data to MongoDB!", err)
	} else {
		fmt.Println("GitHub data saved successfully! ")
	}

	done <- true
	return err
}
