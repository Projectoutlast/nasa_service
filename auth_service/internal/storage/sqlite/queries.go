package sqlite

func (s *SQLiteStorage) addDefaultServicesForUser(email string) error {

	stmt := `INSERT INTO allowed_services (email) VALUES (?)`

	_, err := s.db.Exec(stmt, email)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	return nil
}
