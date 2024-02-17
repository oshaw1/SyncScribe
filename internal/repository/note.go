package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/oshaw1/SyncScribe/internal/model"
)

type NoteRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

// Adjust the function signature to expect a DynamoDB client and a table name as string
func NewNoteRepository(dynamoDBClient *dynamodb.DynamoDB, tableName string) *NoteRepository {
	return &NoteRepository{
		db:        dynamoDBClient,
		tableName: tableName,
	}
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
