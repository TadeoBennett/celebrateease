package postgresql

import (
	"database/sql"
	"tadeobennett/celebrateease/model"
)

type UserModel struct {
	DB *sql.DB
}

// returns all users from the database
func (m *UserModel) GetUsersFromDB() ([]*model.User, error) {
	s := `
	SELECT first_name, last_name from users;
	`
	rows, err := m.DB.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*model.User{}

	for rows.Next() {

		record := &model.User{}

		err = rows.Scan(&record.FirstName, &record.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
