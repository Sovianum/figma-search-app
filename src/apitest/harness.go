package apitest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davyzhang/agw"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Ctx context.Context

	Module *APIModule

	Handler agw.GatewayHandler
}

func (s *Suite) SetupSuite() {
	s.Ctx = context.Background()

	s.Module = s.newModule()

	s.dropAllTables()
	s.createAllTables()

	r := s.Module.API.NewRouter()

	s.Handler = s.createTestHandler(r)
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func (s *Suite) CallEndpoint(path string, body []byte) Response {
	req := &events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       path,
		Body:       string(body),
	}

	bReq, err := json.Marshal(req)
	s.Require().NoError(err)

	result, err := s.Handler(s.Ctx, bReq)
	s.Require().NoError(err)

	resultJSON, err := json.Marshal(result)
	s.Require().NoError(err)

	var resp Response
	s.Require().NoError(json.Unmarshal(resultJSON, &resp))

	return resp
}

func (s *Suite) createAllTables() {
	s.createTagsTable()
}

func (s *Suite) createTagsTable() {
	s.createTableSync(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("pId"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeB),
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeB),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("pId"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("tags"),
	})
}

func (s *Suite) createTableSync(input *dynamodb.CreateTableInput) {
	_, err := s.Module.DynamoDB.CreateTable(input) // todo sync
	s.Require().NoError(err)
}

func (s *Suite) dropAllTables() {
	s.Module.DynamoDB.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String("tags"),
	})
}

func (s *Suite) newModule() *APIModule {
	module, err := InitializeModule()
	s.Require().NoError(err)

	return module
}

func (s *Suite) createTestHandler(h http.Handler) agw.GatewayHandler {
	return func(ctx context.Context, content json.RawMessage) (interface{}, error) {
		agp := agw.NewAPIGateParser(content)
		return agw.Process(agp, h), nil
	}
}
