package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oshaw1/SyncScribe/internal/handler"
)

func main() {
	r := gin.Default()

	// Initialize handlers
	noteHandler := handler.NewNoteHandler()

	// Setup routes
	r.POST("/notes", noteHandler.CreateNote)
	r.GET("/notes/:id", noteHandler.GetNote)
	r.GET("/notes", noteHandler.GetAllNotes)
	r.PUT("/notes/:id", noteHandler.UpdateNote)
	r.DELETE("/notes/:id", noteHandler.DeleteNote)

	// Start server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
