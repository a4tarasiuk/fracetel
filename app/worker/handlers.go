package worker

import (
	"encoding/json"
	"log"

	"fracetel/core/telemetry"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func registerCarTelemetryConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.CarTelemetryTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			carTelemetryMessage := telemetry.CarTelemetry{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &carTelemetryMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			carTelemetry := carTelemetryFromMessage(carTelemetryMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", telemetry.CarTelemetryTopicName, carTelemetry)

			insertToCollection(carTelemetry, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerLapDataConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.LapDataTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			lapDataMessage := telemetry.LapData{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &lapDataMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			lapData := lapDataFromMessage(lapDataMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", telemetry.LapDataTopicName, lapData)

			insertToCollection(lapData, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerSessionConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.SessionTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			sessionMessage := telemetry.Session{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &sessionMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			session := sessionFromMessage(sessionMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", telemetry.SessionTopicName, session)

			insertToCollection(session, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerCarStatusConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.CarStatusTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			carStatusMessage := telemetry.CarStatus{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &carStatusMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			carStatus := carStatusFromMessage(carStatusMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", telemetry.CarStatusTopicName, carStatus)

			insertToCollection(carStatus, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerCarDamageConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.CarDamageTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			carDamageMessage := telemetry.CarDamage{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &carDamageMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			carDamage := carDamageFromMessage(carDamageMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", telemetry.CarDamageTopicName, carDamage)

			insertToCollection(carDamage, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerSessionHistoryConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.SessionHistoryTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			sessionHistoryMessage := telemetry.SessionHistory{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &sessionHistoryMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			sessionHistory := sessionHistoryFromMessage(sessionHistoryMessage, message.Header)

			log.Printf("received msg with [%s] subject: %+v", telemetry.SessionHistoryTopicName, sessionHistory)

			insertToCollection(sessionHistory, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}

func registerFinalClassificationConsumer(natsConn *nats.Conn, collection *mongo.Collection) {
	_, err := natsConn.Subscribe(
		telemetry.FinalClassificationTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			finalClassificationMessage := telemetry.FinalClassification{}

			message := telemetry.Message{
				Header:  telemetry.Header{},
				Payload: &finalClassificationMessage,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			finalClassification := finalClassificationFromMessage(finalClassificationMessage, message.Header)

			log.Printf(
				"received msg with [%s] subject: %+v",
				telemetry.FinalClassificationTopicName,
				finalClassification,
			)

			insertToCollection(finalClassification, collection)
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
