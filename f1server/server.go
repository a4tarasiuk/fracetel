package f1server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"fracetel/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type f1UDPServer struct {
	addr net.IP
	port int

	messageChannel chan *models.Message

	rabbitChannel *amqp.Channel
}

func NewF1UDPServer(
	addr net.IP,
	port int,
	rabbitChannel *amqp.Channel,
) *f1UDPServer {
	return &f1UDPServer{
		addr:          addr,
		port:          port,
		rabbitChannel: rabbitChannel,
	}
}

func CreateAndStart(
	addr net.IP,
	port int,
	rabbitMQURL string,
) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"fracetel_logs",
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare an exchange", err)
	}

	server := NewF1UDPServer(addr, port, ch)

	server.Start()
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
			publishMessage(s.rabbitChannel, message)
		}
	}
}

func publishMessage(ch *amqp.Channel, message interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := json.Marshal(message)
	failOnError(err, "Failed to marshal message to JSON")

	err = ch.PublishWithContext(
		ctx,
		"fracetel_logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	failOnError(err, "Failed to send message to RabbitMQ")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
