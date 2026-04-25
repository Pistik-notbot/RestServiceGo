package main

import (
	"context"
	"fmt"

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

	_, err = simple_sql.SelectRows(conn, ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("succeed")
}
