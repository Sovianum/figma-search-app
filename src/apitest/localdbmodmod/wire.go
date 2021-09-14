package localdbmodmod

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/wire"
)

var M = wire.NewSet(
	NewDB,
)

func NewDB() (dynamodbiface.DynamoDBAPI, error) {
	sessionConfig := aws.NewConfig().
		WithRegion("eu-central-1").
		WithEndpoint("http://localhost:8000")

	dbSession, err := session.NewSession(sessionConfig)
	if err != nil {
		return nil, err
	}

	return dynamodb.New(dbSession), nil
}
