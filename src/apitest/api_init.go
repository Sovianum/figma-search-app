//+build wireinject

package apitest

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiwire"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/wire"
)

func InitializeAPI() *api.API {
	wire.Build(wire.NewSet(
		apiwire.M,
		NewDynamoDB,
	))

	return &api.API{}
}

func NewDynamoDB() dynamodbiface.DynamoDBAPI {
	return nil
}
