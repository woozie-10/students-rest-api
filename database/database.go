package database

import (
	"context"
	"fmt"
	"github.com/woozie-10/students-rest-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var StudentsCollection *mongo.Collection

// InitDB initializes the database connection and verifies the connection to MongoDB.
func InitDB() error {
	url := config.Config.GetString("db.url")
	clientOptions := options.Client().ApplyURI(url)

	var err error

	client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	if err := client.Database("Students").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}

	StudentsCollection = client.Database(config.Config.GetString("db.database")).Collection("studentsCollection")

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}
