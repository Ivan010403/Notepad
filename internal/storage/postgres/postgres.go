package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

//TODO: make all vars from env file or config

const (
	host     = "localhost"
	port     = 8089
	user     = "postgres"
	password = "vany2003"
	dbname   = "Notepad"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	const op = "internal.storage.postgresql.New"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveNote(author string, topic string, content string) error {
	const op = "internal.storage.postgresql.SaveArticle"

	stmt := "INSERT INTO articles(author, topic, content, date) VALUES ($1, $2, $3, $4)"
	curr := time.Now()

	//TODO: may be many errors!!! try to fix error checking

	_, err := s.db.Exec(stmt, author, topic, content, fmt.Sprint(curr.Date()))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

//TODO: make a rules for deleting (you can delete only your own articles)

func (s *Storage) DeleteNote(id string) error {
	const op = "internal.storage.postgresql.DeleteNote"

	stmt := "DELETE FROM articles WHERE id=$1"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) UpdateNote(id int, content string) error {
	const op = "internal.storage.postgresql.DeleteNote"

	stmt := "UPDATE articles SET content=$1 WHERE id=$2"
	_, err := s.db.Exec(stmt, content, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
