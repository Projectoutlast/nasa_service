package main

import (
	"database/sql"
	"os"

	"github.com/Projectoutlast/space_service/auth_service/internal/app"
	"github.com/Projectoutlast/space_service/auth_service/internal/config"
	"github.com/Projectoutlast/space_service/auth_service/internal/logging"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	newLogger := logging.New(cfg.Environment, os.Stdout)

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	application := app.New(newLogger, cfg, db)

	application.GRPCServer.MustRun()
}
