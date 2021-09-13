package tagimpl

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func (dao *dao) FindProjectTags(ctx context.Context, projectID projectid.ID) ([]*tag.Tag, error) {
	filter := expression.Name("pId").Equal(expression.Value(projectID))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return nil, err
	}

	output, err := dao.db.ScanWithContext(ctx, &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	})
	if err != nil {
		return nil, err
	}

	if len(output.LastEvaluatedKey) > 0 {
		panic(errorx.AssertionFailed.New("got unprocessed keys on tags extraction %s", spew.Sdump(output)))
	}

	result := make([]*tag.Tag, 0, len(output.Items))
	for _, respItem := range output.Items {
		var resultItem Tag
		if err := dynamodbattribute.UnmarshalMap(respItem, &resultItem); err != nil {
			return nil, err
		}

		result = append(result, &tag.Tag{
			ID:   resultItem.ID,
			Text: resultItem.Text,
		})
	}

	return result, nil
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

type tagRemovalRequestProj struct {
	ID        tagid.ID     `dynamodbav:"id"`
	ProjectID projectid.ID `dynamodbav:"pId"`
}

func (dao *dao) DeleteTags(ctx context.Context, projectID projectid.ID, tagIDs []tagid.ID) error {
	requests := make([]*dynamodb.WriteRequest, 0, len(tagIDs))
	for _, id := range tagIDs {
		valueMap, err := dynamodbattribute.MarshalMap(tagRemovalRequestProj{
			ID:        id,
			ProjectID: projectID,
		})
		if err != nil {
			return err
		}

		requests = append(requests, &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: valueMap,
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
