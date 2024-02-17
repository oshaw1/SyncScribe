package service

import (
	"time"

	"github.com/oshaw1/SyncScribe/internal/model"
	"github.com/oshaw1/SyncScribe/internal/repository"
)

type NoteService struct {
	noteRepository *repository.NoteRepository
}

// Adjust the NewNoteService function to only require a NoteRepository
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

	err := s.noteRepository.Create(note) // This method should be implemented in the repository
	if err != nil {
		return err
	}

	return nil
}

func (s *NoteService) GetNoteByID(id string) (*model.Note, error) {
	note, err := s.noteRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return note, nil
}
