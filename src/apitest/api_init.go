//+build wireinject

package apitest

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiwire"
	"github.com/Sovianum/figma-search-app/src/apitest/localdbmodmod"
	"github.com/google/wire"
)

func InitializeAPI() (*api.API, error) {
	panic(wire.Build(M))
}

var M = wire.NewSet(
	apiwire.M,
	localdbmodmod.M,
)
