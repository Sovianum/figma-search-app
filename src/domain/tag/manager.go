package tag

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
)

type Manager interface {
	GetTags(ctx context.Context, projectID projectid.ID) ([]*Tag, error)
	CreateTags(ctx context.Context, projectID projectid.ID, tags []*Tag) error
	RemoveTags(ctx context.Context, projectID projectid.ID, tagIDs []tagid.ID) error
}
