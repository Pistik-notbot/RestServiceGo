package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertTask(conn *pgx.Conn, ctx context.Context) error {
	sqlQuery := `
	Insert into tasks (title, description, completed, created_at)
	values ('Task 1', 'Description for task 1', false, now());
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
