package f1tel

import (
	"log"
	"net"

	"fracetel/app/f1tel/packets"
	"fracetel/core/messages"
)

const BufferSizeBytes = 2048

type telemetryServer struct {
	addr net.IP
	port int

	messageStream MessagePublisher
}

func NewTelemetryServer(
	addr net.IP,
	port int,
	messageStream MessagePublisher,
) *telemetryServer {
	return &telemetryServer{
		addr:          addr,
		port:          port,
		messageStream: messageStream,
	}
}

func (s *telemetryServer) StartAndListen() {
	addr := net.UDPAddr{
		IP:   s.addr,
		Port: s.port,
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Fatalf("Failed to listen UDP: %v", err)
	}

	defer conn.Close()

	log.Printf("Listening on %d", s.port)

	messageChan := make(chan *messages.Message)

	go MessageProcessor(s.messageStream, messageChan)

	for {
		buffer := make([]byte, BufferSizeBytes)

		nRead, _, err := conn.ReadFrom(buffer)

		if err != nil {
			log.Printf("Error during reading packets: %v\n", err)
		}

		rawPacket := buffer[:nRead]

		header, err := packets.ParserPacketHeader(rawPacket)
		if err != nil {
			log.Printf("Error during reading Message: %s", err)
			continue
		}

		packetID := packets.ID(header.PacketID)

		parser, err := packets.GetParserForPacket(packetID)
		if err != nil {
			continue
		}

		message, err := parser.ToMessage(header, rawPacket)
		if err != nil {
			continue
		}

		messageChan <- message
	}
}
