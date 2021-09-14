package tagmod

import (
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagimpl"
	"github.com/google/wire"
)

var M = wire.NewSet(
	tagimpl.NewDAO,
	tagimpl.NewManager,
	tagimpl.NewTagger,
)
