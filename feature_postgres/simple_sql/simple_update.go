package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateTask(conn *pgx.Conn, ctx context.Context) error {
	sqlQuery := `
	Update tasks
	Set completed = true, compeleted_at = now()
	Where id = 1;
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err

}
