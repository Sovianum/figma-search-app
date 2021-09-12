package tagimpl

import (
	"context"

	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
)

func NewManager(
	dao *dao,
) tag.Manager {

	return &manager{
		dao: dao,
	}
}

type manager struct {
	dao *dao
}

var _ tag.Manager = (*manager)(nil)

func (m *manager) GetTags(ctx context.Context, projectID projectid.ID) ([]*tag.Tag, error) {
	panic("implement me")
}

func (m *manager) CreateTags(ctx context.Context, projectID projectid.ID, tags []*tag.Tag) error {
	return m.dao.InsertTags(ctx, projectID, tags)
}

func (m *manager) RemoveTags(ctx context.Context, projectID projectid.ID, tagIDs []tagid.ID) error {
	panic("implement me")
}
