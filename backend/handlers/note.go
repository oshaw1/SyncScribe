package handlers

import (
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/google/uuid"
    "log"
    "net/http"
    "github.com/oshaw1/SyncScribe/backend/SyncScribe/db"
    "github.com/oshaw1/SyncScribe/backend/SyncScribe/models"
    "time"
)

func HandleNotes(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getNotes(w, r)
    case "POST":
        createNote(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getNotes(w http.ResponseWriter, r *http.Request) {
    result, err := db.DB.Scan(&dynamodb.ScanInput{
        TableName: aws.String("notes"),
    })
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    var notes []models.Note
    err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &notes)
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(notes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
    var note models.Note
    err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    note.NoteID = uuid.New().String()
    note.CreatedAt = time.Now().Format(time.RFC3339)
    note.UpdatedAt = note.CreatedAt

    _, err = db.DB.PutItem(&dynamodb.PutItemInput{
        TableName: aws.String("notes"),
        Item: map[string]*dynamodb.AttributeValue{
            "NoteID":    {S: aws.String(note.NoteID)},
            "CreatedAt": {S: aws.String(note.CreatedAt)},
            "Content":   {S: aws.String(note.Content)},
            "Tags":      {SS: aws.StringSlice(note.Tags)},
            "Title":     {S: aws.String(note.Title)},
            "UpdatedAt": {S: aws.String(note.UpdatedAt)},
            "UserID":    {S: aws.String(note.UserID)},
        },
    })
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(note)
}