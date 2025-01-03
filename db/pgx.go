package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pgx-contrib/pgxotel"
)

const otelServiceName = "example-pgx"

func NewPGXPool(connString string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal(err)
	}

	// add tracing to the connection
	config.ConnConfig.Tracer = &pgxotel.QueryTracer{
		Name: otelServiceName,
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}

func NewPGXConn(connString string) *pgx.Conn {
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		log.Fatal(err)
	}

	config.Tracer = &pgxotel.QueryTracer{
		Name: otelServiceName,
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
