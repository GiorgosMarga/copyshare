package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Create(title, content string, userId, expires int) (int, error) {
	stmt := `INSERT INTO snippet (title, content, expiresAt,createdAt,userId)
				VALUES (?, ?, DATE_ADD( UTC_TIMESTAMP(), INTERVAL ? DAY), UTC_TIMESTAMP(),?);`
	res, err := s.DB.Exec(stmt, title, content, expires, userId)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *SnippetModel) GetSnippet(id, userId int) (*Snippet, error) {
	sn := &Snippet{}
	stmt := "SELECT id,title,content,createdAt,expiresAt FROM snippet WHERE expiresAt > UTC_TIMESTAMP() AND id = ? AND userId = ?"
	err := s.DB.QueryRow(stmt, id, userId).Scan(&sn.ID, &sn.Title, &sn.Content, &sn.CreatedAt, &sn.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecond
		}
		return nil, err

	}
	return sn, nil

}
func (s *SnippetModel) GetLatestSnippets(userId int) ([]*Snippet, error) {
	stmt := `SELECT id,title,content,createdAt,expiresAt FROM snippet WHERE userId = ? AND  expiresAt > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`
	rows, err := s.DB.Query(stmt, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := []*Snippet{}
	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	// After rows.Scan call rows.Err() to retrieve any error
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil

}
func (s *SnippetModel) DeleteSnippet(id, userId int) error {
	var exists bool
	stmt := `DELETE FROM snippet WHERE id = ? AND expiresAt > UTC_TIMESTAMP() AND userId=?`
	existsStmt := `SELECT EXISTS(SELECT id FROM snippet WHERE id = ? AND expiresAt > UTC_TIMESTAMP() AND userId=?)`
	err := s.DB.QueryRow(existsStmt, id, userId).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return ErrNoRecond
	}

	_, err = s.DB.Exec(stmt, id, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoRecond
		}
		return err
	}
	return nil
}
