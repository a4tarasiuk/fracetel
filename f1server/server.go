package f1server

import (
	"log"
	"net"

	"fracetel/models"
)

type f1UDPServer struct {
	addr net.IP
	port int

	messageChannel chan *models.Message
}

func NewF1UDPServer(
	addr net.IP,
	port int,
	messageChannel chan *models.Message,
) *f1UDPServer {
	return &f1UDPServer{
		addr:           addr,
		port:           port,
		messageChannel: messageChannel,
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

	for {
		buffer := make([]byte, 2048)

		nRead, _, err := conn.ReadFrom(buffer)

		if err != nil {
			log.Printf("Error during reading packets: %v\n", err)
		}

		message, _ := ParsePacket(buffer[:nRead])

		if message != nil {
			s.messageChannel <- message // TODO: Replace blocking call with putting message in a queue
		}
	}
}
