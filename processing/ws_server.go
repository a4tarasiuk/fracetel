package processing

import (
	"encoding/json"
	"log"
	"net/http"

	"fracetel/models"
	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader

	messageCh <-chan *models.Message
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsh.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := wsh.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	defer c.Close()

	for {
		for msg := range wsh.messageCh {

			rawData, err := json.Marshal(msg)

			if err != nil {
				log.Printf("Error %s when marshaling message", err)
			}

			err = c.WriteMessage(websocket.TextMessage, rawData)

			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
		}
	}
}

func StartWsServer(wsMessageCh <-chan *models.Message) {
	wsHandler := webSocketHandler{
		upgrader:  websocket.Upgrader{},
		messageCh: wsMessageCh,
	}

	http.Handle("/", wsHandler)

	log.Print("Starting server...")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
