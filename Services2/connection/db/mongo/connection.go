package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectionToMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("[MongoDB] Connection Failed: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("[MongoDB] Failed to ping the database: ", err)
	}

	fmt.Println("[MongoDB] Connection Successful")
}
