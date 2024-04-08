package postgresql

import (
	"database/sql"
	"tadeobennett/celebrateease/model"
)

type PageModel struct {
	DB *sql.DB
}

// returns all page from the database
func (m *PageModel) GetPagesFromDB() ([]*model.Page, error) {
	s := `
	SELECT * from pages;
	`
	rows, err := m.DB.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	page := []*model.Page{}

	for rows.Next() {
		record := &model.Page{}
		err = rows.Scan(&record.ID, &record.Page_creator_id, &record.Event_id, &record.Title, &record.HeaderImage, &record.ContentImage1, &record.VoiceNote, &record.BodyTextContent, &record.ViewCode, &record.CongratsGIF, &record.LoveGIF, &record.Status, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}
		page = append(page, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return page, nil
}
