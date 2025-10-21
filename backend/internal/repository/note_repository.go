package repository

import (
	"database/sql"

	"github.com/lucasmsaluno/my-notes/internal/model"
)

type NoteRepository interface {
	GetAll() ([]model.Note, error)
	Create(content string) (model.Note, error)
	Update(id int, content string) error
	Delete(id int) error
}

type SQLiteNoteRepo struct {
	db *sql.DB
}

func NewSQLiteNoteRepo(db *sql.DB) *SQLiteNoteRepo {
	return &SQLiteNoteRepo{db}
}

func (r *SQLiteNoteRepo) GetAll() ([]model.Note, error) {
	rows, err := r.db.Query("SELECT id, content FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *SQLiteNoteRepo) Create(content string) (model.Note, error) {
	result, err := r.db.Exec("INSERT INTO notes (content) VALUES (?)", content)
	if err != nil {
		return model.Note{}, err
	}
	id, _ := result.LastInsertId()
	return model.Note{ID: int(id), Content: content}, nil
}

func (r *SQLiteNoteRepo) Update(id int, content string) error {
	_, err := r.db.Exec("UPDATE notes SET content = ? WHERE id = ?", content, id)
	return err
}

func (r *SQLiteNoteRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id = ?", id)
	return err
}
