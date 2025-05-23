package web

import (
	"log"
	"net/http"

	"fracetel/pkg/telemetry"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader

	natsConn      *nats.Conn
	telemetryChan <-chan *telemetry.Message
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsh.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := wsh.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	defer conn.Close()

	_, err = wsh.natsConn.Subscribe(
		telemetry.FRaceTelTopicName, func(msg *nats.Msg) {

			if err = conn.WriteMessage(websocket.TextMessage, msg.Data); err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
		},
	)

	if err != nil {
		log.Printf("failed to subscribe to nats core from ws server: %s", err)
	}

	for {
		messageType, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if messageType == websocket.CloseMessage {
			return
		}
	}
}

func StartWsServerAndListen(natsConn *nats.Conn) {
	telemetryChan := make(chan *telemetry.Message)

	wsHandler := webSocketHandler{
		upgrader:      websocket.Upgrader{},
		natsConn:      natsConn,
		telemetryChan: telemetryChan,
	}

	http.Handle("/", wsHandler)

	log.Print("Starting server...")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
