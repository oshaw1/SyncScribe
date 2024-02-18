package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/oshaw1/SyncScribe/internal/handler"
	"github.com/oshaw1/SyncScribe/internal/repository"
	"github.com/oshaw1/SyncScribe/internal/service"
	"github.com/oshaw1/SyncScribe/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	// Initialize AWS session with region from config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(cfg.AWSRegion),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoDBClient := dynamodb.New(sess)

	// Initialize repository with DynamoDB client and table name from config
	noteRepository := repository.NewNoteRepository(dynamoDBClient, cfg.DynamoDBTable)

	// Initialize service with repository
	noteService := service.NewNoteService(noteRepository)

	// Initialize handlers with service
	noteHandler := handler.NewNoteHandler(noteService)

	r := gin.Default()

	// Setup routes
	r.POST("/notes/CreateNote", noteHandler.CreateNote)
	r.GET("/notes/:id", noteHandler.GetNote)
	r.GET("/notes", noteHandler.GetAllNotesBasedOnUserID)
	r.PUT("/notes/:id", noteHandler.UpdateNote)
	r.DELETE("/notes/DeleteNote/:id", noteHandler.DeleteNote)

	// Start server with port from config
	r.Run("localhost:" + cfg.ServerPort) // listen and serve on configured port
}
