package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertTask(conn *pgx.Conn, ctx context.Context, task TaskModel) error {
	sqlQuery := `
	Insert into tasks (title, description, completed, created_at)
	values ($1, $2, $3, $4);
	`

	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt)

	return err
}
