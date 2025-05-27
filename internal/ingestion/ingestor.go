package ingestion

import (
	"context"

	"fracetel/internal/ingestion/consumers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

func ConsumeTelemetryMessages(ctx context.Context, natsConn *nats.Conn, db *pgxpool.Pool) {
	consumers.RegisterLapData(ctx, natsConn, db)
	consumers.RegisterFinalClassification(ctx, natsConn, db)
	consumers.RegisterSession(ctx, natsConn, db)
	consumers.RegisterSessionHistory(ctx, natsConn, db)
}
