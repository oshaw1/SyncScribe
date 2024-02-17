package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/oshaw1/SyncScribe/internal/model"
)

type NoteRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func NewNoteRepository(dynamoDBClient *dynamodb.DynamoDB, tableName string) *NoteRepository {
	return &NoteRepository{
		db:        dynamoDBClient,
		tableName: tableName,
	}
}

func (r *NoteRepository) Create(note model.Note) error {
	av, err := dynamodbattribute.MarshalMap(note)
	if err != nil {
		fmt.Println("Got error marshalling new note item:", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.tableName),
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:", err)
		return err
	}

	fmt.Println("Successfully added note to DynamoDB table")
	return nil
}
func (r *NoteRepository) Delete(noteID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"NoteID": {
				S: aws.String(noteID),
			},
		},
	}

	_, err := r.db.DeleteItem(input)
	if err != nil {
		fmt.Printf("Got error calling DeleteItem: %s\n", err)
		return err
	}

	fmt.Printf("Successfully deleted note with ID %s from DynamoDB table\n", noteID)
	return nil
}
