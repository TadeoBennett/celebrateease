package postgresql

import (
	"errors"
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
// gets a user from the database with a provided id
func (m *UserModel) GetUserFromDBWithId(id int) (*model.User, error) {
	s := `
	SELECT ID, FirstName, LastName, ProfilePicture, PhoneNumber, DateOfBirth, Email, Password, Status, CreatedAt, UpdatedAt from users WHERE ID = ?;
	`
	record := &model.User{}

	err := m.DB.QueryRow(s, id).Scan(&record.ID,&record.FirstName,&record.LastName,&record.ProfilePicture,&record.PhoneNumber,&record.DateOfBirth,&record.Email,&record.Password,&record.Status,&record.CreatedAt,&record.UpdatedAt)

	if err != nil {
		//allows us to ask go if this error was a specific kind of error
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrRecordNotFound
		}else{
			return nil, err
		}
	}

	// return the row and no error
	return record, nil
}
