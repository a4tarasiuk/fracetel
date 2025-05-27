package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"fracetel/internal/infra"
	"fracetel/internal/ingestion"
	"fracetel/internal/messaging"
	"fracetel/internal/udp"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
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

	natsConn, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS %s", err)
	}

	natsEventStream := messaging.NewNatsEventStream(natsConn)

	f1TelemetryServer := udp.NewTelemetryServer(net.IPv4(0, 0, 0, 0), cfg.F1TelServerPort, natsEventStream)

	go f1TelemetryServer.StartAndListen()

	pgPool, err := pgxpool.New(context.Background(), "postgresql://postgres:postgres@localhost:5432/fracetel")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pgPool.Close()

	ingestion.ConsumeTelemetryMessages(ctx, natsConn, pgPool)

	select {
	case <-ctx.Done():
		stop()
	}

	// TODO: Shutdown app
}
