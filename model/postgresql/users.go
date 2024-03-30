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

		u := &model.User{}

		err = rows.Scan(&u.FirstName, &u.LastName)
		if err != nil {
			return nil, err
		}

		//Append to quotes
		users = append(users, u)
	}

	//Always check the rows.Err()
	//instead of checking for errors and seeing if it is nil, we do it in one line
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
