package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Sovianum/figma-search-app/src/client/clientnode"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-central-1"))

func getTags(userID string, nodeID string) (*clientnode.Node, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"User": {
				S: aws.String(userID),
			},
			"NodeId": {
				S: aws.String(nodeID),
			},
		},
		TableName: aws.String("NodeTags"),
	}

	result, err := db.GetItem(input)
	if err != nil || result.Item == nil {
		return nil, err
	}

	var node clientnode.Node
	if err := dynamodbattribute.UnmarshalMap(result.Item, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

func show(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	node, err := getTags("qwerty", "0:1")
	if err != nil {
		return serverError(err)
	}

	js, err := json.Marshal(node)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func main() {
	lambda.Start(show)
}
