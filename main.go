package main

import (
	"net"

	"fracetel/f1server"
)

func main() {
	f1UDPServer := f1server.NewF1UDPServer(net.IPv4(0, 0, 0, 0), 20777)

	f1UDPServer.Start()
}
