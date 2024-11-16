package main

import (
	"log"
	"net"

	"fracetel/packet"
)

func main() {
	addr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20777,
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Fatalf("Failed to listen UDP: %v", err)
	}

	defer conn.Close()

	log.Println("Listening on 20777")

	for {
		buffer := make([]byte, 2048)

		nRead, _, err := conn.ReadFrom(buffer)

		if err != nil {
			log.Printf("Error during reading packet: %v\n", err)
		}

		// log.Printf("Received packet from: %s, bytes=%d\n", receiveAddr, nRead)

		packet.ParsePacket(buffer[:nRead])
	}
}
