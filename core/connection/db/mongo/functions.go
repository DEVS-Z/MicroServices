package mongo

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Collection(collection string) *mongo.Collection {
	if client == nil {
		log.Fatal("[MongoDB] Client is not connected")
	}
	/**

	Validate in the future if the database and collection are found

	*/
	return client.Database(os.Getenv("MONGO_DB")).Collection(collection)
}

func Schema[T any](collection string) *mongo.Collection {
	return Collection(collection)
}

func CreateIndex(collection string, field string, unique bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := Collection(collection)

	indexModel := mongo.IndexModel{
		Keys: bson.M{field: 1},
		Options: options.Index().
			SetUnique(unique),
	}

	_, err := col.Indexes().CreateOne(ctx, indexModel)

	if mongo.IsDuplicateKeyError(err) {
		return errors.New(field + " already exists")
	}

	return err
}
