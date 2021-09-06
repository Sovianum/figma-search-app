package clienttag

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/tag"
)

type Converter struct{}

func (c *Converter) ConvertTags(ctx context.Context, tags []*tag.Tag) []*Tag {
	result := make([]*Tag, 0, len(tags))
	for _, t := range tags {
		result = append(result, &Tag{
			ID:   t.ID,
			Text: t.Text,
		})
	}

	return result
}
