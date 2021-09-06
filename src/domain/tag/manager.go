package tag

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/files/fileid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
)

type Manager interface {
	GetTags(ctx context.Context, fileID fileid.ID) ([]*Tag, error)
	CreateTags(ctx context.Context, fileID fileid.ID, tags []*Tag) error
	RemoveTags(ctx context.Context, fileID fileid.ID, tagIDs []tagid.ID) error
}
