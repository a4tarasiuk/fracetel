package main

import (
	"log"
	"net"

	"fracetel/app/f1server"
	"fracetel/app/web"
	"fracetel/app/worker"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	foreverCh := make(chan int)

	// TODO: Move RabbitMQ channel init to ./internal/rabbitmq package.
	// 	Create Infra struct and encapsulate all infrastructure initialization there

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	telemetryPublisher := f1server.NewTelemetryMessagePublisher(ch)

	eventPublisher := f1server.NewEventPublisher(ch)

	go web.StartWsServer("amqp://guest:guest@localhost:5672/")

	f1UDPServer := f1server.NewF1UDPServer(net.IPv4(0, 0, 0, 0), 20777, telemetryPublisher, eventPublisher)

	go f1UDPServer.Start()

	go worker.ConsumeEvents(ch)

	<-foreverCh
}
