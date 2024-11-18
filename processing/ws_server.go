package processing

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader

	rabbitChannel *amqp.Channel
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsh.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := wsh.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	defer c.Close()

	q, err := wsh.rabbitChannel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	err = wsh.rabbitChannel.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		"fracetel_logs", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to create a queue binding", err)
	}

	messages, err := wsh.rabbitChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to register a consumer", err)
	}

	for {
		for msg := range messages {

			err = c.WriteMessage(websocket.TextMessage, msg.Body)

			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
		}
	}
}

func StartWsServer(rabbitMQURL string) {
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

	wsHandler := webSocketHandler{
		upgrader:      websocket.Upgrader{},
		rabbitChannel: ch,
	}

	http.Handle("/", wsHandler)

	log.Print("Starting server...")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
