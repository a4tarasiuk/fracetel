package ingestion

import (
	"context"

	"fracetel/internal/ingestion/consumers"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

func ConsumeTelemetryMessages(ctx context.Context, natsConn *nats.Conn, db *pgx.Conn) {
	consumers.RegisterLapData(ctx, natsConn, db)
}
