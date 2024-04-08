package postgresql

import (
	"database/sql"
	"tadeobennett/celebrateease/model"
)

type EventModel struct {
	DB *sql.DB
}

// returns all event from the database
func (m *EventModel) GetEventsFromDB() ([]*model.Event, error) {
	s := `
	SELECT * from events;
	`
	rows, err := m.DB.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	event := []*model.Event{}

	for rows.Next() {
		record := &model.Event{}
		err = rows.Scan(&record.ID, &record.EventName, &record.EventPicture, &record.EventDescription, &record.Status, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}
		event = append(event, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return event, nil
}
