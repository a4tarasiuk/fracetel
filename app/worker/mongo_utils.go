package worker

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func insertToCollection(obj any, collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, obj)

	if err != nil {
		log.Printf("failed to insert obj to collection: %s", err)
	}
}
