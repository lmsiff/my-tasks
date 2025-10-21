package service

import (
	"github.com/lucasmsaluno/my-notes/internal/model"
	"github.com/lucasmsaluno/my-notes/internal/repository"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{repo}
}

func (s *NoteService) GetAllNotes() ([]model.Note, error) {
	return s.repo.GetAll()
}

func (s *NoteService) CreateNote(content string) (model.Note, error) {
	return s.repo.Create(content)
}

func (s *NoteService) UpdateNote(id int, content string) error {
	return s.repo.Update(id, content)
}

func (s *NoteService) DeleteNote(id int) error {
	return s.repo.Delete(id)
}
