//+build wireinject

package apitest

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiwire"
	"github.com/Sovianum/figma-search-app/src/apitest/dbmockmod"
	"github.com/google/wire"
	"github.com/gusaul/go-dynamock"
)

func InitializeAPI() *api.API {
	panic(wire.Build(M))
}

func InitializeMock() *dynamock.DynaMock {
	panic(wire.Build(M))
}

var M = wire.NewSet(
	apiwire.M,
	dbmockmod.M,
)
