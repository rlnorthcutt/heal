package models

import (
	"github.com/rlnorthcutt/heal/internal/db"
)

type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetAllContent() ([]Content, error) {
	rows, err := db.DB.Query("SELECT id, title, body FROM content")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contentList []Content
	for rows.Next() {
		var content Content
		if err := rows.Scan(&content.ID, &content.Title, &content.Body); err != nil {
			return nil, err
		}
		contentList = append(contentList, content)
	}
	return contentList, nil
}

func InsertContent(content Content) error {
	_, err := db.DB.Exec("INSERT INTO content (title, body) VALUES (?, ?)", content.Title, content.Body)
	return err
}
