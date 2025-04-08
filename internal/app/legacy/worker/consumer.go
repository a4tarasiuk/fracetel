package worker

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ConsumeEvents(natsConn *nats.Conn, mongoDB *mongo.Database) {
	registerCarTelemetryConsumer(natsConn, mongoDB.Collection("car_telemetry"))

	registerCarStatusConsumer(natsConn, mongoDB.Collection("car_status"))

	registerCarDamageConsumer(natsConn, mongoDB.Collection("car_damage"))

	registerLapDataConsumer(natsConn, mongoDB.Collection("lap_data"))

	registerSessionConsumer(natsConn, mongoDB.Collection("session"))

	registerSessionHistoryConsumer(natsConn, mongoDB.Collection("session_history"))

	registerFinalClassificationConsumer(natsConn, mongoDB.Collection("final_classification"))
}
