package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oshaw1/SyncScribe/internal/model"
	"github.com/oshaw1/SyncScribe/internal/service"
)

type NoteHandler struct {
	noteService *service.NoteService
}

func NewNoteHandler(s *service.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: s,
	}
}
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateNote without passing the note object, since it's not designed to accept parameters.
	err := h.noteService.CreateNote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Since CreateNote doesn't return the created note, you can't directly return the created note to the client.
	// This is a significant limitation of the current design.
	// You might return a generic success message or the static note details if applicable.
	c.JSON(http.StatusOK, gin.H{"message": "Note created successfully"})
}

func (h *NoteHandler) GetNote(c *gin.Context) {

}
func (h *NoteHandler) GetAllNotesBasedOnUserID(c *gin.Context) {

}
func (h *NoteHandler) UpdateNote(c *gin.Context) {

}
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	// Extracting the note ID from the request URL parameter
	noteID := c.Param("id")
	if noteID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Note ID is required"})
		return
	}

	// Calling the service layer to delete the note by ID
	err := h.noteService.DeleteNoteByID(noteID)
	if err != nil {
		// For simplicity, we're returning an Internal Server Error for all errors I will adapt this
		// and return a 404 Not Found error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If no error occurred, the note was successfully deleted
	c.JSON(http.StatusOK, gin.H{"message": "Note successfully deleted"})
}
