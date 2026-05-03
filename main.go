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

	tasks, err := simple_sql.SelectRows(conn, ctx)
	if err != nil {
		panic(err)
	}

	task := tasks[0]
	task.Completed = true

	err = simple_sql.UpdateTask(conn, ctx, task)
	if err != nil {
		panic(err)
	}

	_, err = simple_sql.SelectRows(conn, ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("succeed")
}
