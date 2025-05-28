package ingestion

import (
	"context"

	"fracetel/internal/infra"
	"fracetel/internal/ingestion/consumers"
)

func ConsumeTelemetryMessages(ctx context.Context, app infra.Infra) {
	consumers.RegisterLapData(ctx, app.NatsConn, app.DB)
	consumers.RegisterFinalClassification(ctx, app.NatsConn, app.DB)
	consumers.RegisterSession(ctx, app.NatsConn, app.DB)
	consumers.RegisterSessionHistory(ctx, app.NatsConn, app.DB)
}
