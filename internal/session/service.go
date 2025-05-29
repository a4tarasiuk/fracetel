package session

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	Create(ctx context.Context, sessionID string) (Session, error)
}

func NewService(db *pgxpool.Pool) Service {
	return postgresSessionService{db: db}
}
