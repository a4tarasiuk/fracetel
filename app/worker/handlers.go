package worker

import (
	"context"
	"encoding/json"
	"log"

	"fracetel/core/messages"
	"fracetel/core/streams"
	"github.com/nats-io/nats.go/jetstream"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func registerCarTelemetryConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	carTelemetryConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "car_telemetry_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.CarTelemetrySubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create carTelemetryConsumer: %s", err)
	}

	_, err = carTelemetryConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			carTelemetryMessage := messages.CarTelemetry{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &carTelemetryMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			carTelemetry := carTelemetryFromMessage(carTelemetryMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", streams.CarTelemetrySubjectName, carTelemetry)

			createCarTelemetryRecord(carTelemetry, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerLapDataConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	lapDataConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "lap_data_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.LapDataSubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create lapDataConsumer: %s", err)
	}

	_, err = lapDataConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			lapDataMessage := messages.LapData{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &lapDataMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			lapData := lapDataFromMessage(lapDataMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", streams.CarTelemetrySubjectName, lapData)

			insertToCollection(lapData, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
