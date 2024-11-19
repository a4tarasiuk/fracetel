package main

import (
	"net"

	"fracetel/app/f1server"
	"fracetel/app/web"
)

func main() {
	foreverCh := make(chan int)

	go web.StartWsServer("amqp://guest:guest@localhost:5672/")

	go f1server.CreateAndStart(
		net.IPv4(0, 0, 0, 0),
		20777,
		"amqp://guest:guest@localhost:5672/",
	)

	<-foreverCh
}
