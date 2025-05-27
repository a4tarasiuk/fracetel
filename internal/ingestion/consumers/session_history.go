package consumers

import (
	"context"
	"encoding/json"
	"log"

	"fracetel/internal/messaging"
	"fracetel/pkg/telemetry"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

func RegisterSessionHistory(ctx context.Context, natsConn *nats.Conn, db *pgxpool.Pool) {
	_, err := natsConn.Subscribe(
		messaging.SessionHistoryTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			sessionHistory := telemetry.SessionHistory{}

			message := messaging.Message{
				Header:  messaging.Header{},
				Payload: &sessionHistory,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			log.Printf("received msg with [%s] subject: %+v", messaging.SessionHistoryTopicName, sessionHistory)

			conn, err := db.Acquire(ctx)
			defer conn.Release()

			if err != nil {
				log.Fatalf("cannot aquire db conn from pool")
			}

			query := `
			INSERT INTO session_history_telemetry (
			                      session_id, 
			                      frame_identifier,
			                      occurred_at, 
			                      num_laps,
								  best_lap_time_lap_num,
								  best_sector_1_lap_num,
								  best_sector_2_lap_num,
								  best_sector_3_lap_num,
								  laps_history
		  	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			`

			rawLapsHistory, err := json.Marshal(sessionHistory.LapsHistory)

			if err != nil {
				log.Fatalf("failed to marhal laps history")
			}

			_, err = conn.Exec(
				context.Background(),
				query,
				sessionHistory.SessionID,
				sessionHistory.FrameIdentifier,
				sessionHistory.OccurredAt,
				sessionHistory.NumLaps,
				sessionHistory.BestLapTimeLapNum,
				sessionHistory.BestSector1LapNum,
				sessionHistory.BestSector2LapNum,
				sessionHistory.BestSector3LapNum,
				rawLapsHistory,
			)

			if err != nil {
				log.Fatalf("failed to insert session history record to db: %s", err)
			}
		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
