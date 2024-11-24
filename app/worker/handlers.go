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

			log.Printf("received msg with [%s] subject: %+v", streams.LapDataSubjectName, lapData)

			insertToCollection(lapData, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerSessionConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	sessionConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "session_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.SessionSubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create sessionConsumer: %s", err)
	}

	_, err = sessionConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			sessionMessage := messages.Session{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &sessionMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			session := sessionFromMessage(sessionMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", streams.SessionSubjectName, session)

			insertToCollection(session, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerCarStatusConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	carStatusConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "car_status_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.CarStatusSubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create carStatusConsumer: %s", err)
	}

	_, err = carStatusConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			carStatusMessage := messages.CarStatus{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &carStatusMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			carStatus := carStatusFromMessage(carStatusMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", streams.CarStatusSubjectName, carStatus)

			insertToCollection(carStatus, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerCarDamageConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	carDamageConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "car_damage_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.CarDamageSubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create carDamageConsumer: %s", err)
	}

	_, err = carDamageConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			carDamageMessage := messages.CarDamage{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &carDamageMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			carDamage := carDamageFromMessage(carDamageMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", streams.CarDamageSubjectName, carDamage)

			insertToCollection(carDamage, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerSessionHistoryConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	sessionHistoryConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "session_history_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.SessionHistorySubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create sessionHistoryConsumer: %s", err)
	}

	_, err = sessionHistoryConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			sessionHistoryMessage := messages.SessionHistory{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &sessionHistoryMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			sessionHistory := sessionHistoryFromMessage(sessionHistoryMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", streams.SessionHistorySubjectName, sessionHistory)

			insertToCollection(sessionHistory, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerFinalClassificationConsumer(js jetstream.JetStream, ctx context.Context, collection *mongo.Collection) {
	finalClassificationConsumer, err := js.CreateConsumer(
		ctx,
		streams.FRaceTelStreamName,
		jetstream.ConsumerConfig{
			Durable:       "final_classification_consumer",
			AckPolicy:     jetstream.AckExplicitPolicy,
			FilterSubject: streams.FinalClassificationSubjectName,
		},
	)
	if err != nil {
		log.Fatalf("failed to create finalClassificationConsumer: %s", err)
	}

	_, err = finalClassificationConsumer.Consume(
		func(jsMsg jetstream.Msg) {
			jsMsg.Ack()

			finalClassificationMessage := messages.FinalClassification{}

			message := messages.Message{
				Header:  messages.Header{},
				Payload: &finalClassificationMessage,
			}

			if err = json.Unmarshal(jsMsg.Data(), &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			finalClassification := finalClassificationFromMessage(finalClassificationMessage, message.Header)

			log.Printf(
				"received msg with [%s] subject: %+v",
				streams.FinalClassificationSubjectName,
				finalClassification,
			)

			insertToCollection(finalClassification, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
