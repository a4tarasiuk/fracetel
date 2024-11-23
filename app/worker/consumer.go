package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"fracetel/core/messages"
	"fracetel/core/streams"
	"github.com/nats-io/nats.go/jetstream"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ConsumeEvents(js jetstream.JetStream, mongoClient *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

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
		log.Fatalf("failed to create sessionStartedConsumer: %s", err)
	}

	carTelemetryCollection := mongoClient.Database("fracetel").Collection("car_telemetries")

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

			createCarTelemetryRecord(carTelemetry, carTelemetryCollection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}

}

//
// func ConsumeEvents(js jetstream.JetStream, userSessionRepository sessions.UserSessionRepository) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
//
// 	sessionStartedConsumer, err := js.CreateConsumer(
// 		ctx,
// 		streams.SessionStreamName,
// 		jetstream.ConsumerConfig{
// 			Durable:       "session_started_consumer",
// 			AckPolicy:     jetstream.AckExplicitPolicy,
// 			FilterSubject: streams.SessionStartedSubject,
// 		},
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to create sessionStartedConsumer: %s", err)
// 	}
//
// 	_, err = sessionStartedConsumer.Consume(
// 		func(msg jetstream.Msg) {
// 			fmt.Printf("Received a JetStream message via sessionStartedConsumer: %s\n", string(msg.Data()))
//
// 			message := messages.Message{}
//
// 			if err = json.Unmarshal(msg.Data(), &message); err != nil {
// 				log.Printf("failed to unmarshal message: %s", err)
// 				return
// 			}
// 			if err = sessionStartedHandler(message, userSessionRepository); err != nil {
// 				log.Printf("failed to persist session: %s", err)
// 			}
//
// 			if err = msg.Ack(); err != nil {
// 				log.Printf("failed to ack message: %s", err)
// 			}
// 		},
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to consume messages err: %s", err)
// 	}
//
// 	sessionFinishedConsumer, err := js.CreateConsumer(
// 		ctx,
// 		streams.SessionStreamName,
// 		jetstream.ConsumerConfig{
// 			Durable:       "session_finished_consumer",
// 			AckPolicy:     jetstream.AckExplicitPolicy,
// 			FilterSubject: streams.SessionFinishedSubject,
// 			// TODO: Cover case when message is failed to process and needs to be retried. How many retries? Interval?
// 		},
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to create sessionStartedConsumer: %s", err)
// 	}
//
// 	_, err = sessionFinishedConsumer.Consume(
// 		func(msg jetstream.Msg) {
// 			fmt.Printf("Received a JetStream message via sessionFinishedConsumer: %s\n", string(msg.Data()))
//
// 			message := messages.Message{}
//
// 			if err = json.Unmarshal(msg.Data(), &message); err != nil {
// 				log.Printf("failed to unmarshal message: %s", err)
// 				return
// 			}
//
// 			if err = sessionFinishedHandler(message, userSessionRepository); err != nil {
// 				log.Printf("failed to update session: %s", err)
// 			}
//
// 			if err = msg.Ack(); err != nil {
// 				log.Printf("failed to ack message: %s", err)
// 			}
// 		},
// 	)
// }
//
// func sessionStartedHandler(message messages.Message, userSessionRepository sessions.UserSessionRepository) error {
// 	session := sessions.UserSession{
// 		ExternalID: strconv.FormatUint(message.SessionID, 10),
// 		StartedAt:  message.OccurredAt,
// 		FinishedAt: nil,
// 	}
//
// 	return userSessionRepository.Create(session)
// }
//
// func sessionFinishedHandler(message messages.Message, userSessionRepository sessions.UserSessionRepository) error {
// 	session, err := userSessionRepository.GetByExternalID(message.SessionID)
//
// 	if err != nil {
// 		log.Printf("Unnable to retrieve session: %s", err)
// 		return sessions.SessionDoesNotExist
// 	}
//
// 	session.FinishedAt = &message.OccurredAt
//
// 	return userSessionRepository.Update(session)
// }
