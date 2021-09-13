//+build wireinject

package apitest

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiwire"
	"github.com/Sovianum/figma-search-app/src/apitest/localdbmodmod"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/wire"
)

func InitializeModule() (*APIModule, error) {
	panic(wire.Build(M))
}

var M = wire.NewSet(
	apiwire.M,
	localdbmodmod.M,
	NewAPIModule,
)

type APIModule struct {
	API *api.API

	DynamoDB dynamodbiface.DynamoDBAPI
}

func NewAPIModule(
	api *api.API,
	db dynamodbiface.DynamoDBAPI,
) *APIModule {
	return &APIModule{
		API:      api,
		DynamoDB: db,
	}
}
