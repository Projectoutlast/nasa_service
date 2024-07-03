package sqlite

import (
	"database/sql"
	"log/slog"
)

type SQLiteStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func New(log *slog.Logger, storagePath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &SQLiteStorage{
		log: log,
		db:  db,
	}, nil
}
