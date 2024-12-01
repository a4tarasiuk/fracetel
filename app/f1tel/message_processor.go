package f1tel

import (
	"log"

	"fracetel/core/messages"
	"fracetel/core/streams"
)

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
