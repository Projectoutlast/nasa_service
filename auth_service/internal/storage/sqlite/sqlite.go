package sqlite

import (
	"database/sql"
	"log/slog"
)

type SQLiteStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func New(log *slog.Logger, db *sql.DB) *SQLiteStorage {

	return &SQLiteStorage{
		log: log,
		db:  db,
	}
}
