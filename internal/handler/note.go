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
	// Extract the `id` parameter from the URL
	id := c.Param("id")

	// Call the service layer to get the note by ID
	note, err := h.noteService.GetNoteByID(id)
	if err != nil {
		// If the note is not found or any other error occurs, return an HTTP error
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	// If the note is found, return it
	c.JSON(http.StatusOK, note)
}
func (h *NoteHandler) GetAllNotes(c *gin.Context) {

}
func (h *NoteHandler) UpdateNote(c *gin.Context) {

}
func (h *NoteHandler) DeleteNote(c *gin.Context) {

}
