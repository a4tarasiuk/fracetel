package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"fracetel/internal/infra"
	"fracetel/internal/ingestion"
	"fracetel/internal/messaging"
	"fracetel/internal/udp"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())

	signalStream := make(chan os.Signal, 1)
	signal.Notify(signalStream, syscall.SIGINT, syscall.SIGTERM)

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
	defer natsConn.Close()

	natsEventStream := messaging.NewNatsEventStream(natsConn)

	udpTelemetryServer := udp.NewTelemetryServer(net.IPv4(0, 0, 0, 0), cfg.F1TelServerPort, natsEventStream)

	go udpTelemetryServer.StartAndListen(ctx)

	pgPool, err := pgxpool.New(context.Background(), "postgresql://postgres:postgres@localhost:5432/fracetel")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pgPool.Close()

	ingestion.ConsumeTelemetryMessages(ctx, natsConn, pgPool)

	<-signalStream

	stop() // TODO: Replace with defer when infra and app abstraction are added
}
