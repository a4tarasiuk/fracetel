package main

import (
	"context"
	"net"
	"time"

	"fracetel/app/f1tel"
	"fracetel/app/worker"
	"fracetel/internal/infra"
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

	js := infra.InitJetStream(ctx)

	mongoClient, _ := mongo.Connect(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	messageStream := f1tel.NewJetStreamMessagePublisher(js)

	f1TelemetryServer := f1tel.NewTelemetryServer(net.IPv4(0, 0, 0, 0), 20777, messageStream)

	go f1TelemetryServer.StartAndListen()

	go worker.ConsumeEvents(js, mongoClient)

	//
	// go web.StartWsServer()

	<-foreverCh
}
