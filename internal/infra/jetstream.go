package infra

import (
	"context"
	"log"

	"fracetel/core/streams"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func InitJetStream(ctx context.Context) jetstream.JetStream {
	nc, _ := nats.Connect(nats.DefaultURL)

	js, _ := jetstream.New(nc)

	// session stream
	_, err := js.CreateStream(
		ctx, jetstream.StreamConfig{
			Name:      streams.SessionStreamName,
			Subjects:  []string{streams.SessionStreamSubject},
			Retention: jetstream.WorkQueuePolicy,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create stream (SessionStreamName): %s", err)
	}

	_, err = js.CreateStream(
		ctx, jetstream.StreamConfig{
			Name:      streams.FRaceTelStreamName,
			Subjects:  []string{streams.FRaceTelSubjectName},
			Retention: jetstream.WorkQueuePolicy,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create stream (FRaceTelStreamName): %s", err)
	}

	return js
}
