package sqlite

import (
	"github.com/Projectoutlast/space_service/auth_service/internal/models/storage"
)

func (s *SQLiteStorage) Registration(email, hashedPassword string) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}

	defer tx.Rollback()

	stmt := `INSERT INTO users (email, password) VALUES (?, ?)`

	res, err := s.db.Exec(stmt, email, hashedPassword)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}

	err = s.addDefaultServicesForUser(email)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		s.log.Error(err.Error())
		return 0, err
	}

	return res.LastInsertId()
}

func (s *SQLiteStorage) GetUser(email string) (*storage.User, error) {
	stmt := `SELECT id, email, password FROM users WHERE email = ?`

	row := s.db.QueryRow(stmt, email)

	var user storage.User
	err := row.Scan(&user.ID, &user.Email, &user.Hash)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return &user, nil
}

func (s *SQLiteStorage) GetUserServices(email string) ([]string, error) {
	stmt := `SELECT services FROM allowed_services WHERE email = ?`

	row := s.db.QueryRow(stmt, email)

	services, err := s.unmarshalledUserServicesSlice(row)
	if err != nil {
		return nil, err
	}

	return services, nil
}
