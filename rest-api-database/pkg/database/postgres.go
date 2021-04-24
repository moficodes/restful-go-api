package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func PGPool(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
