package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, "CONN_STRING")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
