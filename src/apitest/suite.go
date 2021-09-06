package apitest

import (
	"context"
	"net/http/httptest"

	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Ctx context.Context

	API *api.API

	server *httptest.Server
}

func (s *Suite) SetupSuite() {
	s.API = s.newAPI()

	r := s.API.NewRouter()

	s.server = httptest.NewServer(r)
}

func (s *Suite) TearDownSuite() {
	r := s.API.NewRouter()

	s.server = httptest.NewServer(r)
}

func (s *Suite) newAPI() *api.API {
	tagEndpoints := api.NewTagEndpoints(nil, nil, nil)

	return api.NewAPI(tagEndpoints)
}
