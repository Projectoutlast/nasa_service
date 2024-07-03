package sqlite

import "github.com/Projectoutlast/space_service/auth_service/internal/models/storage"

func (s *SQLiteStorage) Registration(email, hashedPassword string) (int64, error) {
	stmt := `INSERT INTO users (email, password) VALUES (?, ?)`

	res, err := s.db.Exec(stmt, email, hashedPassword)
	if err != nil {
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
	stmt := `SELECT services FROM users WHERE email = ?`

	row := s.db.QueryRow(stmt, email)

	var services []string
	err := row.Scan(&services)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return services, nil
}
