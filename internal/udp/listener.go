package udp

import (
	"log"
	"net"

	"fracetel/internal/messaging"
	"fracetel/internal/udp/packets"
	"fracetel/pkg/telemetry"
)

const BufferSizeBytes = 2048

type telemetryServer struct {
	addr net.IP
	port int

	eventStream messaging.EventStream
}

func NewTelemetryServer(
	addr net.IP,
	port int,
	eventStream messaging.EventStream,
) *telemetryServer {
	return &telemetryServer{
		addr:        addr,
		port:        port,
		eventStream: eventStream,
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

	telMessageChan := make(chan *telemetry.Message, 100)

	go messagePublisher(s.eventStream, telMessageChan)

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

		telMessage, err := parser.ToTelemetryMessage(header, rawPacket)
		if err != nil {
			continue
		}

		telMessageChan <- telMessage
	}
}
