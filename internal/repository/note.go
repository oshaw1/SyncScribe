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

func (r *NoteRepository) FindByID(id string) (*model.Note, error) {
	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Notes"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var note model.Note
	err = dynamodbattribute.UnmarshalMap(result.Item, &note)
	if err != nil {
		return nil, err
	}

	return &note, nil
}
