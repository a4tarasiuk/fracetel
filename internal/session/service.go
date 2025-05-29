package session

import (
	"context"

	"fracetel/internal/infra"
)

type Service interface {
	Create(ctx context.Context, sessionID string) (Session, error)
}

func NewService(infra infra.Infra) Service {
	return postgresSessionService{infra.DB}
}
