package worker

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func createCarTelemetryRecord(telemetry CarTelemetry, collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, telemetry)

	if err != nil {
		log.Printf("failed to insert car telemetry to collection: %s", err)
	}
}
