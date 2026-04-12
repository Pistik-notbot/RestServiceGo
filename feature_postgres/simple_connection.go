package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
}
