package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
