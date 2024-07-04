package test_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Projectoutlast/space_service/auth_service/internal/logging"
	"github.com/Projectoutlast/space_service/auth_service/internal/storage/sqlite"
	"github.com/stretchr/testify/require"
)

type user struct {
	email    string
	password string
}

func TestStorage(t *testing.T) {
	storage, db := initStorage(t)

	createTable(t, db)

	newUser := user{
		email:    "test@example.com",
		password: "12345678",
	}

	user_id, err := storage.Registration(newUser.email, newUser.password)
	require.NoError(t, err)
	require.IsType(t, int64(10), user_id)

	existsUser, err := storage.GetUser(newUser.email)
	require.NoError(t, err)
	require.Equal(t, newUser.email, existsUser.Email)

	userService, err := storage.GetUserServices(existsUser.Email)
	require.NoError(t, err)
	require.Equal(t, userService[0], "nasa_random")

	db.Close()
	_, err = storage.Registration(newUser.email, newUser.password)
	require.Error(t, err)
}

func initStorage(t *testing.T) (*sqlite.SQLiteStorage, *sql.DB) {
	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)

	newLogger := logging.New("development", os.Stdout)

	storage := sqlite.New(newLogger, db)

	return storage, db
}

func createTable(t *testing.T, db *sql.DB) {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL);`,
		`CREATE TABLE IF NOT EXISTS allowed_services (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, services ARRAY NOT NULL DEFAULT '["nasa_random"]', FOREIGN KEY (email) REFERENCES users(email));`,
	}

	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			require.NoError(t, err)
		}
	}
}
