package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connString string) (*Database, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	db := Database{pool: pool}
	err = db.pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database %w", err)
	}
	return &db, nil
}

func (db *Database) Pool() *pgxpool.Pool {
	return db.pool
}

func (db *Database) Close() {
	db.pool.Close()
}
