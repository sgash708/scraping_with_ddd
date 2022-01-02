package repository

import (
	"context"
	"database/sql"
)

type OfficeRepository interface {
	Insert(ctx context.Context, db *sql.DB, datas map[string]map[string]string, isForeign, tableName string) int
}
