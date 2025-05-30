package ingestion

import (
	"context"
	"encoding/json"
	"log"

	"fracetel/internal/messaging"
	"fracetel/internal/session"
	"fracetel/pkg/telemetry"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

func RegisterFinalClassification(ctx context.Context, natsConn *nats.Conn, db *pgxpool.Pool) {
	_, err := natsConn.Subscribe(
		messaging.FinalClassificationTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			finalClassification := telemetry.FinalClassification{}

			message := messaging.Message{
				Header:  messaging.Header{},
				Payload: &finalClassification,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			log.Printf(
				"Received msg with [%s] subject: %+v",
				messaging.FinalClassificationTopicName,
				finalClassification,
			)

			conn, err := db.Acquire(ctx)
			defer conn.Release()

			if err != nil {
				log.Fatalf("cannot aquire db conn from pool")
			}

			query := `
			INSERT INTO final_classifications (
			                      session_id, 
			                      frame_identifier,
			                      occurred_at, 
			                      finishing_position,
							   	  starting_position
		  	) VALUES ($1, $2, $3, $4, $5)
			`

			_, err = conn.Exec(
				context.Background(),
				query,
				finalClassification.SessionID,
				finalClassification.FrameIdentifier,
				finalClassification.OccurredAt,
				finalClassification.FinishingPosition,
				finalClassification.StartingPosition,
			)

			if err != nil {
				log.Fatalf("failed to insert final classification record to db: %s", err)
			}

			go func() {
				_, err = session.NewService(db).Create(ctx, finalClassification.SessionID)
				if err != nil {
					log.Printf("Failed to create session: %s", err)
				}
			}()

		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
