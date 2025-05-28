package udp

import (
	"context"
	"log"
	"net"

	"fracetel/internal/messaging"
	"fracetel/internal/udp/packets"
)

const BufferSizeBytes = 2048

type telemetryServer struct {
	addr net.IP
	port int

	eventStream messaging.EventStream
}

func NewTelemetryServer(
	port int,
	eventStream messaging.EventStream,
) *telemetryServer {
	return &telemetryServer{
		addr:        net.IPv4(0, 0, 0, 0),
		port:        port,
		eventStream: eventStream,
	}
}

func (s *telemetryServer) StartAndListen(ctx context.Context) {
	addr := net.UDPAddr{
		IP:   s.addr,
		Port: s.port,
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Fatalf("Failed to listen UDP: %v", err)
	}

	telemetryMessageChan := make(chan *messaging.Message, 100)

	defer func() {
		conn.Close()
		close(telemetryMessageChan)
	}()

	log.Printf("UDP server is listening on :%d", s.port)

	go messagePublisher(s.eventStream, telemetryMessageChan)

	go func() {
		for {
			buffer := make([]byte, BufferSizeBytes)

			nRead, _, err := conn.ReadFrom(buffer)

			if err != nil {
				log.Printf("Error during reading packets: %v\n", err)
				return
			}

			rawPacket := buffer[:nRead]

			header, err := packets.ParserPacketHeader(rawPacket)
			if err != nil {
				log.Printf("Error during reading Message: %s", err)
				continue
			}

			parser, err := packets.GetParserForPacket(packets.ID(header.PacketID))
			if err != nil {
				continue
			}

			telemetryMessage, err := parser.ToTelemetryMessage(header, rawPacket)
			if err != nil {
				continue
			}

			telemetryMessageChan <- telemetryMessage
		}
	}()

	<-ctx.Done()

	log.Println("UDP server is shutting down")
}
