package apitest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/aws/aws-lambda-go/events"
	"github.com/davyzhang/agw"
	"github.com/gusaul/go-dynamock"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Ctx context.Context

	API    *api.API
	DBMock *dynamock.DynaMock

	Handler agw.GatewayHandler
}

func (s *Suite) SetupSuite() {
	s.Ctx = context.Background()

	s.API, s.DBMock = s.newAPI()

	r := s.API.NewRouter()

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

func (s *Suite) newAPI() (*api.API, *dynamock.DynaMock) {
	testAPI := InitializeAPI()
	return testAPI.API, testAPI.Mock
}

func (s *Suite) createTestHandler(h http.Handler) agw.GatewayHandler {
	return func(ctx context.Context, content json.RawMessage) (interface{}, error) {
		agp := agw.NewAPIGateParser(content)
		return agw.Process(agp, h), nil
	}
}
