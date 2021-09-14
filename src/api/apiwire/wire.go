package apiwire

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/client/clienttag/clienttagmod"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagmod"
	"github.com/google/wire"
)

var M = wire.NewSet(
	api.NewAPI,
	api.NewTagEndpoints,

	tagmod.M,
	clienttagmod.M,
)
