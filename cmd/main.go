package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"

	"fracetel/internal/app/legacy/web"
	"fracetel/internal/app/legacy/worker"
	"fracetel/internal/infra"
	"fracetel/internal/messaging"
	"fracetel/internal/udp"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// TODO:
	//  1. Create module interface that accepts all the infra staff (natsConn, mongo.Database)
	//  2. Add implementation that creates needed interface implementations with provided infra objects
	//  3. Allow to start-up the sub-applications via module interface
	//  The main goal is to encapsulate infra staff and module logic with components

	// TODO: Add graceful shutdown

	// 	ctx, cancel := context.WithCancel(context.Background())
	//	ch := make(chan os.Signal, 1)
	//	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	// 	<-ch
	//	cancel()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	oTelShutdown, err := infra.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, oTelShutdown(context.Background()))
	}()

	cfg := infra.LoadConfigFromEnv()

	mongoClient, _ := mongo.Connect(options.Client().ApplyURI(cfg.DBUrl))
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mongoDB := mongoClient.Database(cfg.DBName)

	natsConn, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS %s", err)
	}

	natsEventStream := messaging.NewNatsEventStream(natsConn)

	f1TelemetryServer := udp.NewTelemetryServer(net.IPv4(0, 0, 0, 0), cfg.F1TelServerPort, natsEventStream)

	go f1TelemetryServer.StartAndListen()

	go worker.ConsumeEvents(natsConn, mongoDB)

	go web.StartWsServerAndListen(natsConn)

	select {
	case <-ctx.Done():
		stop()
	}

	// TODO: Shutdown app
}
