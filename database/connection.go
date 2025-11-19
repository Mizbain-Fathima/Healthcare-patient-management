package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// var PatientColl *mongo.Collection

func ConnectDB() {
	godotenv.Load()
	uri := os.Getenv("ATLAS_CONNECTION_STRING")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	// PatientColl = Client.Database("health_db").Collection("patients")

	log.Println("Connected to MongoDB Atlas!")
}

func GetPatientCollection() *mongo.Collection {
	return Client.Database("health_db").Collection("patients")
}
