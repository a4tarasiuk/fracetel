package main

import (
	"net"

	"fracetel/f1server"
	"fracetel/processing"
)

func main() {
	foreverCh := make(chan int)

	go processing.StartWsServer("amqp://guest:guest@localhost:5672/")

	go f1server.CreateAndStart(
		net.IPv4(0, 0, 0, 0),
		20777,
		"amqp://guest:guest@localhost:5672/",
	)

	<-foreverCh
}
