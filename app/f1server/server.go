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
}

func NewF1UDPServer(
	addr net.IP,
	port int,
	messageStream MessagePublisher,
) *f1UDPServer {
	return &f1UDPServer{
		addr:          addr,
		port:          port,
		messageStream: messageStream,
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

	messageChan := make(chan *messages.Message)

	go MessageProcessor(s.messageStream, messageChan)

	for {
		buffer := make([]byte, 2048)

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

func MessageProcessor(messageStream MessagePublisher, messageChan <-chan *messages.Message) {
	for message := range messageChan {

		subjectName, ok := streams.MessageTypeSubjectMap[message.Type]

		if !ok {
			continue
		}

		if err := messageStream.Publish(message, subjectName); err != nil {
			log.Printf("failed to publish message. type: %s| packed_id: %s", message.Type, message.Header.PacketID)
		}
	}
}
