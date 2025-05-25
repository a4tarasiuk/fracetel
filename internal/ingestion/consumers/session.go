package consumers

import (
	"context"
	"encoding/json"
	"log"

	"fracetel/internal/messaging"
	"fracetel/pkg/telemetry"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

func RegisterSession(ctx context.Context, natsConn *nats.Conn, db *pgx.Conn) {
	_, err := natsConn.Subscribe(
		messaging.SessionTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			session := telemetry.Session{}

			message := messaging.Message{
				Header:  messaging.Header{},
				Payload: &session,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			log.Printf("received msg with [%s] subject: %+v", messaging.SessionTopicName, session)

			query := `
			INSERT INTO session_telemetry (
			                      session_id, 
			                      frame_identifier,
			                      occurred_at, 
			                      weather, 
			                      track_temperature, 
			                      air_temperature, 
			                      total_laps, 
			                      track_length, 
			                      track_id, 
								  type,
			                      time_left, 
			                      duration
		  	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			`

			_, err := db.Exec(
				context.Background(),
				query,
				session.SessionID,
				session.FrameIdentifier,
				session.OccurredAt,
				session.Weather,
				session.TrackTemperature,
				session.AirTemperature,
				session.TotalLaps,
				session.TrackLength,
				session.TrackID,
				session.Type,
				session.TimeLeft,
				session.Duration,
			)

			if err != nil {
				log.Fatalf("failed to insert session record to db: %s", err)
			}

		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
