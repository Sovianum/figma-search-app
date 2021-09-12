//+build wireinject

package apitest

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiwire"
	"github.com/Sovianum/figma-search-app/src/apitest/dbmockmod"
	"github.com/google/wire"
	"github.com/gusaul/go-dynamock"
)

func InitializeAPI() *TestAPI {
	panic(wire.Build(M))
}

var M = wire.NewSet(
	NewTestAPI,
	apiwire.M,
	dbmockmod.M,
)

func NewTestAPI(api *api.API, mock *dynamock.DynaMock) *TestAPI {
	return &TestAPI{
		API:  api,
		Mock: mock,
	}
}

type TestAPI struct {
	API  *api.API
	Mock *dynamock.DynaMock
}
