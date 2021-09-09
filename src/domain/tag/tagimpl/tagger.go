package tagimpl

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/tag"
)

func NewTagger() tag.Tagger {
	return &tagger{}
}

type tagger struct{}

var _ tag.Tagger = (*tagger)(nil)

func (t *tagger) TagNodes(ctx context.Context, query tag.TagNodesQuery) error {
	panic("implement me")
}
