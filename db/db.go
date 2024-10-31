package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() *mongo.Client {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var MongoDBUsername = os.Getenv("MONGO_DB_USERNAME")
	var MongoDBPassword = os.Getenv("MONGO_DB_PASSWORD")
	var MongoDBURI = fmt.Sprintf("mongodb+srv://%s:%s@monfernape.lvp08.mongodb.net/?retryWrites=true&w=majority&appName=Monfernape", MongoDBUsername, MongoDBPassword)
	fmt.Println(MongoDBURI)
	fmt.Println(MongoDBUsername)
	fmt.Println(MongoDBPassword)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoDBURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}
