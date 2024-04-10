package postgresql

import (
	"database/sql"
	"tadeobennett/celebrateease/model"
)

type CelebrantModel struct {
	DB *sql.DB
}

// returns all celebrant from the database
func (m *CelebrantModel) GetCelebrantsFromDB() ([]*model.Celebrant, error) {
	s := `
	SELECT * FROM celebrants;
	`
	rows, err := m.DB.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	celebrant := []*model.Celebrant{}

	for rows.Next() {
		record := &model.Celebrant{}
		err = rows.Scan(&record.ID, &record.Creator_id, &record.FirstName, &record.LastName, &record.Email, &record.DateOfBirth, &record.Status, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}
		celebrant = append(celebrant, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return celebrant, nil
}
// returns all celebrant from the database to be loaded on the dashboard home page
func (m *CelebrantModel) GetCelebrantsForLoggedInUserFromDB(id int) ([]*model.Celebrant, error) {
	s := `
	SELECT * FROM Celebrants AS C
	WHERE C.Creator_id = ?;
	`
	rows, err := m.DB.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	celebrant := []*model.Celebrant{}

	for rows.Next() {
		record := &model.Celebrant{}
		err = rows.Scan(&record.ID, &record.Creator_id, &record.FirstName, &record.LastName, &record.Email, &record.DateOfBirth, &record.Status, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}
		celebrant = append(celebrant, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return celebrant, nil
}
