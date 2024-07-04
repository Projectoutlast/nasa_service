package sqlite

import (
	"database/sql"
	"encoding/json"
)

func (s *SQLiteStorage) unmarshalledUserServicesSlice(row *sql.Row) ([]string, error) {
	var servicesJSON string
	if err := row.Scan(&servicesJSON); err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	var services []string
	if err := json.Unmarshal([]byte(servicesJSON), &services); err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return services, nil
}
