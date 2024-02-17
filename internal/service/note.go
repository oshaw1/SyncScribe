package service

import (
	"time"

	"github.com/oshaw1/SyncScribe/internal/model"
	"github.com/oshaw1/SyncScribe/internal/repository"
)

type NoteService struct {
	noteRepository *repository.NoteRepository
}

func NewNoteService(noteRepo *repository.NoteRepository) *NoteService {
	return &NoteService{
		noteRepository: noteRepo,
	}
}

func (s *NoteService) CreateNote() error {
	note := model.Note{
		NoteID:    "2",
		CreatedAt: time.Now().Format(time.RFC3339),
		Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Tags:      []string{},
		Title:     "TestNote2",
		UpdatedAt: time.Now().Format(time.RFC3339),
		UserID:    "1",
	}

	err := s.noteRepository.Create(note)
	if err != nil {
		return err
	}

	return nil
}

func (s *NoteService) DeleteNoteByID(noteID string) error {
	return s.noteRepository.Delete(noteID)
}
