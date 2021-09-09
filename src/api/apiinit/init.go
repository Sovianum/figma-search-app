//+build wireinject

package apiinit

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/api/apiwire"
	"github.com/google/wire"
)

func InitializeAPI() *api.API {
	wire.Build(apiwire.M)

	return &api.API{}
}
