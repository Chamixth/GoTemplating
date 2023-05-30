package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Document represents a document entity
type Document struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// MongoDB configuration
const (
	timeout = 5 * time.Second
	letter  = `
Username:  {{.Username}},
Email: {{.Email}},
Password: {{.Password}}
`
)

// MongoDB client and collection instances
var (
	client *mongo.Client
	coll   *mongo.Collection
)

// Initialize initializes the MongoDB client and collection
func Initialize(uri, dbName, collection string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	coll = client.Database(dbName).Collection(collection)
	return nil
}

// RetrieveDocuments retrieves all documents from the MongoDB collection
func RetrieveDocuments() ([]Document, error) {
	// Access the desired collection
	collection := client.Database(coll.Database().Name()).Collection(coll.Name())

	// Define an empty slice to store the documents
	var documents []Document

	// Prepare the find options (if needed)
	findOptions := options.Find()

	// Execute the find query
	cur, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Iterate over the results
	for cur.Next(context.Background()) {
		var document Document

		err := cur.Decode(&document)
		if err != nil {
			return nil, err
		}

		// Append the document to the slice
		documents = append(documents, document)
	}

	// Check for any cursor errors
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}

func main() {
	// MongoDB configuration
	uri := "mongodb+srv://chamith:123@cluster0.ujlq82i.mongodb.net/?retryWrites=true&w=majority"
	dbName := "sample_db2"
	collection := "Users2"

	err := Initialize(uri, dbName, collection)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// Define a handler to retrieve and render the documents
	e.GET("/documents", func(c echo.Context) error {
		documents, err := RetrieveDocuments()
		if err != nil {
			return err
		}

		t := template.Must(template.New("letter").Parse(letter))

		for _, r := range documents {
			err := t.Execute(os.Stdout, r)
			if err != nil {
				log.Println("Executing template", err)
			}
		}

		return c.String(http.StatusOK, "Documents retrieved")
	})

	// Start the server
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Close the MongoDB connection when the server stops
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
