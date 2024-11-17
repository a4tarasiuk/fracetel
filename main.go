package main

import (
	"net"

	"fracetel/f1server"
	"fracetel/models"
	"fracetel/processing"
)

func main() {

	wsMessageCh := make(chan *models.Message)

	messageCh := make(chan *models.Message)

	consumer := processing.NewMessageConsumer(messageCh, wsMessageCh)

	go consumer.ProcessMessages()

	go processing.StartWsServer(wsMessageCh)

	f1UDPServer := f1server.NewF1UDPServer(net.IPv4(0, 0, 0, 0), 20777, messageCh)

	f1UDPServer.Start()
}
