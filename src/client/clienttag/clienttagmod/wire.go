package clienttagmod

import (
	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/google/wire"
)

var M = wire.NewSet(clienttag.NewConverter)
