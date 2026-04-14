package main

import (
	"context"

	simple_connection "go.mod/feature_postgres/simple_connection"
	"go.mod/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	simple_sql.CreateTable(conn, ctx)
}
