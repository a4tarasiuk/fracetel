package consumers

import (
	"context"

	"fracetel/internal/consumers/ingestion"
	"fracetel/internal/infra"
)

func Register(ctx context.Context, app infra.Infra) {
	ingestion.RegisterLapData(ctx, app.NatsConn, app.DB)
	ingestion.RegisterFinalClassification(ctx, app.NatsConn, app.DB)
	ingestion.RegisterSession(ctx, app.NatsConn, app.DB)
	ingestion.RegisterSessionHistory(ctx, app.NatsConn, app.DB)
}
