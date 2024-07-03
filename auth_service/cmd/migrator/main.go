package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	// Драйвер для выполнения миграций SQLite 3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// Драйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "путь к хранилищу")

	flag.StringVar(&migrationsPath, "migrations-path", "", "путь к папке с миграциями")

	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "имя таблицы с миграциями")

	flag.Parse()

	if storagePath == "" || migrationsPath == "" {
		panic("не указан путь к хранилищу или папке с миграциями")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Print("нечего обновлять")

			return
		}

		panic(err)
	}
}
