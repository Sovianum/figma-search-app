package apitest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiinit"
	"github.com/davyzhang/agw"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Ctx context.Context

	API *api.API

	handler agw.GatewayHandler
}

func (s *Suite) SetupSuite() {
	s.Ctx = context.Background()

	s.API = s.newAPI()

	r := s.API.NewRouter()

	s.handler = s.createTestHandler(r)
}

func (s *Suite) newAPI() *api.API {
	return apiinit.InitializeAPI()
}

func (s *Suite) createTestHandler(h http.Handler) agw.GatewayHandler {
	return func(ctx context.Context, content json.RawMessage) (interface{}, error) {
		agp := agw.NewAPIGateParser(content)
		return agw.Process(agp, h), nil
	}
}
