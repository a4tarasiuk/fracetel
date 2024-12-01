package main

import (
	"context"
	"log"
	"net"

	"fracetel/app/f1tel"
	"fracetel/app/web"
	"fracetel/internal/infra"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// TODO:
	//  1. Create module interface that accepts all the infra staff (natsConn, mongo.Database)
	//  2. Add implementation that creates needed interface implementations with provided infra objects
	//  3. Allow to start-up the sub-applications via module interface
	//  The main goal is to encapsulate infra staff and module logic with componenets

	foreverCh := make(chan int)

	mongoClient, _ := mongo.Connect(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS %s", err)
	}

	natsEventStream := infra.NewNatsEventStream(natsConn)

	f1TelemetryServer := f1tel.NewTelemetryServer(net.IPv4(0, 0, 0, 0), 20777, natsEventStream)

	go f1TelemetryServer.StartAndListen()

	// go worker.ConsumeEvents(js, mongoClient)

	go web.StartWsServerAndListen(natsConn)

	<-foreverCh
}
