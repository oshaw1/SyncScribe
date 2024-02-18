package service

import (
	"fmt"
	"math/rand"
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

func (s *NoteService) CreateNote(note *model.Note) error {
	// Now the service layer takes responsibility for setting these fields.
	note.NoteID = generateNoteID()
	now := time.Now().Format(time.RFC3339)
	note.CreatedAt = now
	note.UpdatedAt = now

	return s.noteRepository.Create(*note)
}

func (s *NoteService) DeleteNoteByID(noteID string) error {
	return s.noteRepository.Delete(noteID)
}

func generateNoteID() string {
	// Seed the random number generator to ensure different results on each run
	rand.Seed(time.Now().UnixNano())
	// Generate a random number in the range [10000, 99999]
	id := rand.Intn(90000) + 10000
	// Convert the integer to a string
	return fmt.Sprintf("%d", id)
}
