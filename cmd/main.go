package main

import (
	"context"
	"log"
	"net"
	"time"

	"fracetel/app/f1server"
	"fracetel/app/worker"
	"fracetel/core/streams"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	foreverCh := make(chan int)

	// TODO: Move JetsStream init to ./internal/rabbitmq package.
	// 	Create Infra struct and encapsulate all infrastructure initialization there

	// go web.StartWsServer("amqp://guest:guest@localhost:5672/")

	// In the `jetstream` package, almost all API calls rely on `context.Context` for timeout/cancellation handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nc, _ := nats.Connect(nats.DefaultURL)

	// Create a JetStream management interface
	js, _ := jetstream.New(nc)

	// session stream
	_, err := js.CreateStream(
		ctx, jetstream.StreamConfig{
			Name:      streams.SessionStreamName,
			Subjects:  []string{streams.SessionStreamSubject},
			Retention: jetstream.WorkQueuePolicy,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create stream (SessionStreamName): %s", err)
	}

	_, err = js.CreateStream(
		ctx, jetstream.StreamConfig{
			Name:      streams.FRaceTelStreamName,
			Subjects:  []string{streams.FRaceTelSubjectName},
			Retention: jetstream.WorkQueuePolicy,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create stream (FRaceTelStreamName): %s", err)
	}

	mongoClient, _ := mongo.Connect(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	messageStream := f1server.NewJetStreamMessagePublisher(js)

	f1UDPServer := f1server.NewF1UDPServer(net.IPv4(0, 0, 0, 0), 20777, messageStream)

	go f1UDPServer.Start()

	go worker.ConsumeEvents(js, mongoClient)

	//
	// go web.StartWsServer()

	<-foreverCh
}
