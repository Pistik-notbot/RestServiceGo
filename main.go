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

	if err := simple_sql.CreateTable(conn, ctx); err != nil {
		panic(err)
	}

	fmt.Println("Таблица создана успешно")
}
