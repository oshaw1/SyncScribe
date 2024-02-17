package service

import (
	"github.com/oshaw1/SyncScribe/internal/model"
	"github.com/oshaw1/SyncScribe/internal/repository"
)

type NoteService struct {
	noteRepository *repository.NoteRepository
}

func NewNoteService(repo *repository.NoteRepository) *NoteService {
	return &NoteService{
		noteRepository: repo,
	}
}

func (s *NoteService) CreateNote(note *model.Note) (*model.Note, error) {
	// Here you would call the repository to save the note to the database
	// For now, we'll just simulate by returning the same note
	return note, nil
}

func (s *NoteService) GetNoteByID(id string) (*model.Note, error) {
	// This is where you'd typically interact with the repository to fetch the note
	// For demonstration, let's assume it returns a dummy note or an error
	note, err := s.noteRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return note, nil
}
