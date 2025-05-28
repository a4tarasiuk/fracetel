package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"fracetel/internal/infra"
	"fracetel/internal/ingestion"
	"fracetel/internal/messaging"
	"fracetel/internal/udp"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())

	signalStream := make(chan os.Signal, 1)
	signal.Notify(signalStream, syscall.SIGINT, syscall.SIGTERM)

	oTelShutdown, err := infra.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, oTelShutdown(context.Background()))
	}()

	infra, err := infra.Init(ctx)

	if err != nil {
		infra.Shutdown()
		log.Fatalf("Failed to start application: %s", err)
	}

	natsEventStream := messaging.NewNatsEventStream(infra.NatsConn)

	udpTelemetryServer := udp.NewTelemetryServer(infra.Config.UDPServerPort, natsEventStream)

	go udpTelemetryServer.StartAndListen(ctx)

	ingestion.ConsumeTelemetryMessages(ctx, infra)

	<-signalStream

	stop() // TODO: Replace with defer when infra and app abstraction are added

	if err = infra.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown application: %s", err)
	}
}
