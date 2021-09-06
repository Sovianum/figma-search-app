package tag

import (
	"context"
)

type Tagger interface {
	TagNodes(ctx context.Context, query TagNodesQuery) error
}
