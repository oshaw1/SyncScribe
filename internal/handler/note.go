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

func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		noteService: service.NewNoteService(),
	}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdNote, err := h.noteService.CreateNote(&note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdNote)
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
