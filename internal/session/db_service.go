package session

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresSessionService struct {
	db *pgxpool.Pool
}

func (s postgresSessionService) Create(ctx context.Context, sessionID string) (Session, error) {

	if s.sessionExists(ctx, sessionID) {
		return s.getSessionByID(ctx, sessionID)
	}

	return s.aggregateSession(ctx, sessionID)
}

func (s postgresSessionService) sessionExists(ctx context.Context, sessionID string) bool {
	var sessionExists bool

	s.db.QueryRow(ctx, selectSessionExistsSQLQuery, sessionID).Scan(&sessionExists)

	return sessionExists
}

func (s postgresSessionService) getSessionByID(ctx context.Context, sessionID string) (Session, error) {
	var session Session

	if err := pgxscan.Get(ctx, s.db, &session, selectSessionSQLQuery, sessionID); err != nil {
		return Session{}, err
	}

	return session, nil
}

func (s postgresSessionService) aggregateSession(ctx context.Context, sessionID string) (Session, error) {
	session := Session{ID: sessionID, CreatedAt: time.Now().UTC()}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return Session{}, err
	}

	if err = tx.QueryRow(ctx, selectFinalClassificationSQLQuery, sessionID).Scan(
		&session.StartingPosition,
		&session.FinishingPosition,
	); err != nil {
		return Session{}, err
	}

	if err = tx.QueryRow(ctx, selectSessionReferencesSQLQuery, sessionID).Scan(
		&session.Weather,
		&session.TotalLaps,
		&session.TrackID,
	); err != nil {
		return Session{}, err
	}

	if err = tx.QueryRow(ctx, selectSessionDurationSQLQuery, sessionID).Scan(&session.Duration); err != nil {
		return Session{}, err
	}

	if err = tx.QueryRow(ctx, selectFastestLapStatsSQLQuery, sessionID).Scan(
		&session.FastestLapNumber,
		&session.FastestLapTimeMs,
		&session.FastestLapSector1Ms,
		&session.FastestLapSector2Ms,
		&session.FastestLapSector3Ms,
	); err != nil {
		return Session{}, err
	}

	_, err = tx.Exec(
		ctx,
		insertSessionSQLQuery,
		session.ID,
		session.FastestLapTimeMs,
		session.FastestLapSector1Ms,
		session.FastestLapSector2Ms,
		session.FastestLapSector3Ms,
		session.FastestLapNumber,
		session.TotalLaps,
		session.TrackID,
		session.Weather,
		session.Duration,
		session.StartingPosition,
		session.FinishingPosition,
		session.CreatedAt,
	)
	if err != nil {
		return Session{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return Session{}, err
	}

	return session, nil
}
