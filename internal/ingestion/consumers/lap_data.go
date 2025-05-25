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

func RegisterLapData(ctx context.Context, natsConn *nats.Conn, db *pgx.Conn) {
	_, err := natsConn.Subscribe(
		messaging.LapDataTopicName, func(natsMsg *nats.Msg) {
			natsMsg.Ack()

			lapData := telemetry.LapData{}

			message := messaging.Message{
				Header:  messaging.Header{},
				Payload: &lapData,
			}

			if err := json.Unmarshal(natsMsg.Data, &message); err != nil {
				log.Printf("failed to unmarshal message: %s", err)
				return
			}

			log.Printf("received msg with [%s] subject: %+v", messaging.LapDataTopicName, lapData)

			query := `
			INSERT INTO lap_data_telemetry (
			                      session_id, 
			                      frame_identifier,
			                      occurred_at, 
			                      last_lap_time_ms, 
			                      current_lap_time_ms, 
			                      first_sector_ms, 
			                      second_sector_ms, 
			                      lap_distance, 
			                      total_distance, 
			                      car_position, 
			                      current_lap_num, 
			                      sector, 
			                      driver_status
		  	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
			`

			_, err := db.Exec(
				context.Background(),
				query,
				lapData.SessionID,
				lapData.FrameIdentifier,
				lapData.OccurredAt,
				lapData.LastLapTimeMs,
				lapData.CurrentLapTimeMs,
				lapData.FirstSectorTimeMs,
				lapData.SecondSectorTimeMs,
				lapData.LapDistance,
				lapData.TotalDistance,
				lapData.CarPosition,
				lapData.CurrentLapNum,
				lapData.Sector,
				lapData.DriverStatus,
			)

			if err != nil {
				log.Fatalf("failed to insert lap data record to db: %s", err)
			}

		},
	)
	if err != nil {
		log.Fatalf("failed to run message consumer: %s", err)
	}
}
