package f1server

import (
	"log"
	"net"

	"fracetel/app/f1server/packets"
	"fracetel/core/messages"
	"fracetel/core/streams"
)

type f1UDPServer struct {
	addr net.IP
	port int

	messageStream MessagePublisher

	sessionManager *_sessionManager
}

func NewF1UDPServer(
	addr net.IP,
	port int,
	messageStream MessagePublisher,
) *f1UDPServer {
	return &f1UDPServer{
		addr:           addr,
		port:           port,
		messageStream:  messageStream,
		sessionManager: newSessionStateManager(messageStream),
	}
}

func (s *f1UDPServer) Start() {
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

	jsCorePublisher := NewJSCoreMessagePublisher()

	for {
		buffer := make([]byte, 2048)

		nRead, _, err := conn.ReadFrom(buffer)

		if err != nil {
			log.Printf("Error during reading packets: %v\n", err)
		}

		rawPacket := buffer[:nRead]

		header, err := packets.ParserPacketHeader(rawPacket)
		if err != nil {
			log.Printf("Error during reading Header: %s", err)
			continue
		}

		s.sessionManager.StartSessionIfNotExist(header.SessionUID)

		packetID := packets.ID(header.PacketID)

		parser, err := packets.GetParserForPacket(packetID)
		if err != nil {
			continue
		}

		message, err := parser.ToMessage(header, rawPacket)
		if err != nil {
			continue
		}

		if message.Type == messages.SessionFinishedMessageType {
			s.sessionManager.FinishSession()
		}

		if message.Type == messages.CarTelemetryMessageType {
			jsCorePublisher.Publish(message, "telemetry.*")
		}

		subject, exists := streams.MessageTypeSubjectMap[message.Type]
		if !exists {
			continue
		}

		if err = s.messageStream.Publish(message, subject); err != nil {
			log.Printf("failed to publish message: %s", err)
		}
	}
}
