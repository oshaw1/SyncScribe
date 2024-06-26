package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Notes    []string           `bson:"notes"`
	Allowed  bool               `bson:"allowed"`
}

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"createdAt"`
	Content   string             `bson:"content"`
	Tags      []string           `bson:"tags"`
	Title     string             `bson:"title"`
	UpdatedAt string             `bson:"updatedAt"`
	UserID    []string           `bson:"userID"`
	FolderID  string             `bson:"folderID"`
}

type Folder struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	UserID         []string           `bson:"userID"`
	ParentFolderID string             `bson:"parentFolderID"`
	ChildFolderIDs []string           `bson:"childFolderIDs"`
	NoteIDs        []string           `bson:"noteIDs"`
}
