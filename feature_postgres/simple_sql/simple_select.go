package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(conn *pgx.Conn, ctx context.Context) ([]TaskModel, error) {
	fmt.Println("start")
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	From tasks
	`

	rows, err := conn.Query(ctx, sqlQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
		PrintTask(task)
	}
	return tasks, err
}

func PrintTask(task TaskModel) {
	fmt.Println("---------------------------")
	fmt.Println(task.ID)
	fmt.Println(task.Title)
	fmt.Println(task.Description)
	fmt.Println(task.Completed)
	fmt.Println(task.CreatedAt)
	fmt.Println(task.CompletedAt)
}
