package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(conn *pgx.Conn, ctx context.Context) {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);
	`
	conn.Exec(ctx, sqlQuery)

}
