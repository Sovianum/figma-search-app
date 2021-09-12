package tagimpl

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/davecgh/go-spew/spew"
	"github.com/joomcode/errorx"
)

const (
	tableName = "tags"
)

func NewDAO(dbAPI dynamodbiface.DynamoDBAPI) *dao {
	return &dao{
		db: dbAPI,
	}
}

type dao struct {
	db dynamodbiface.DynamoDBAPI
}

func (dao *dao) InsertTags(ctx context.Context, projectID projectid.ID, tags []*tag.Tag) error {
	requests := make([]*dynamodb.WriteRequest, 0, len(tags))
	for _, t := range tags {
		valueMap, err := dynamodbattribute.MarshalMap(Tag{
			ID:        t.ID,
			Text:      t.Text,
			ProjectID: projectID,
		})
		if err != nil {
			return err
		}

		requests = append(requests, &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: valueMap,
			},
		})
	}

	output, err := dao.db.BatchWriteItemWithContext(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			tableName: requests,
		},
	})
	if err != nil || output == nil {
		return err
	}

	if len(output.UnprocessedItems) > 0 {
		panic(errorx.Panic(errorx.IllegalState.New("unexpected unprocessed items %s", spew.Sdump(output)))) // todo handle it in a sensible way
	}

	return nil
}
