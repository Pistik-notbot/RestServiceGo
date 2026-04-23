package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertTask(
	conn *pgx.Conn,
	ctx context.Context,
	title string,
	description string,
	completed bool,
	created_at time.Time,
) error {
	sqlQuery := `
	Insert into tasks (title, description, completed, created_at)
	values ($1, $2, $3, $4);
	`

	_, err := conn.Exec(ctx, sqlQuery, title, description, completed, created_at)

	return err
}
