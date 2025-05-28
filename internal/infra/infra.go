package infra

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type Infra struct {
	DB *pgxpool.Pool

	NatsConn *nats.Conn

	Config Config
}

func Init(ctx context.Context) (Infra, error) {
	i := Infra{}

	i.Config = LoadConfigFromEnv()

	pgPool, err := pgxpool.New(ctx, i.Config.DBUrl)

	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return Infra{}, err
	}

	i.DB = pgPool

	natsConn, err := nats.Connect(i.Config.NatsURL)
	if err != nil {
		log.Printf("Failed to connect to NATS %s", err)
		return Infra{}, err
	}

	i.NatsConn = natsConn

	return i, nil
}

func (i Infra) Shutdown() error {
	if i.DB != nil {
		i.DB.Close()
	}

	if i.NatsConn != nil {
		i.NatsConn.Close()
	}

	return nil
}
