package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Note represents the note model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Pin      string             `bson:"pin"`
	Notes    []string           `bson:"notes"`
	Allowed  []string           `bson:"allowed"`
}

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	NoteID    string             `bson:"noteID"`
	CreatedAt string             `bson:"createdAt"`
	Content   string             `bson:"content"`
	Tags      []string           `bson:"tags"`
	Title     string             `bson:"title"`
	UpdatedAt string             `bson:"updatedAt"`
	UserID    string             `bson:"userID"`
	FolderID  string             `bson:"folderID"`
}

type Folder struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	UserID         string             `bson:"userID"`
	ParentFolderID string             `bson:"parentFolderID"`
	ChildFolderIDs []string           `bson:"childFolderIDs"`
	NoteIDs        []string           `bson:"noteIDs"`
}
