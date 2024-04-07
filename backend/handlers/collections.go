package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection
var notesCollection *mongo.Collection
var foldersCollection *mongo.Collection

func SetCollections(users, notes, folders *mongo.Collection) {
	usersCollection = users
	notesCollection = notes
	foldersCollection = folders
}

func GetUsersCollection() *mongo.Collection {
	return usersCollection
}

func GetNotesCollection() *mongo.Collection {
	return notesCollection
}

func GetFoldersCollection() *mongo.Collection {
	return foldersCollection
}
