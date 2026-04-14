package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	return conn, nil
}
