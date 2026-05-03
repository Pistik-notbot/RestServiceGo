package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateTask(conn *pgx.Conn, ctx context.Context, task TaskModel) error {
	sqlQuery := `
	Update tasks
	Set title = $1, description = $2, completed = $3, created_at = $4, completed_at = $5
	Where id = $6;
	`
	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt, task.CompletedAt, task.ID)
	return err

}
