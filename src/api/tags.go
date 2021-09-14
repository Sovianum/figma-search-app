package api

import (
	"context"
	"net/http"

	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
	"github.com/Sovianum/figma-search-app/src/url"
	"github.com/Sovianum/figma-search-app/src/util"
)

func NewTagEndpoints(
	tagManager tag.Manager,
	tagConverter *clienttag.Converter,
	tagger tag.Tagger,
) *TagEndpoints {

	return &TagEndpoints{
		tagManager:   tagManager,
		tagConverter: tagConverter,
		tagger:       tagger,
	}
}

type TagEndpoints struct {
	tagManager   tag.Manager
	tagConverter *clienttag.Converter

	tagger tag.Tagger
}

func (ep *TagEndpoints) GetTags(r *http.Request) (interface{}, error) {
	ctx := r.Context()

	projectID, err := url.ProjectIDFromRequest(r)
	if err != nil {
		return nil, err
	}

	return ep.doGetTags(ctx, projectID)
}

type tagCreationRequest struct {
	Tags []*clienttag.Tag `json:"tags"`
}

func (ep *TagEndpoints) CreateTags(r *http.Request) (interface{}, error) {
	ctx := r.Context()

	projectID, err := url.ProjectIDFromRequest(r)
	if err != nil {
		return nil, err
	}

	var req tagCreationRequest
	if err := util.UnmarshalFromReaderCloser(r.Body, &req); err != nil {
		return nil, err
	}

	if err := ep.tagManager.CreateTags(ctx, projectID, ep.toDomainTags(req.Tags)); err != nil {
		return nil, err
	}

	return ep.doGetTags(ctx, projectID)
}

type tagsRemovalRequest struct {
	IDs []tagid.ID `json:"ids"`
}

func (ep *TagEndpoints) RemoveTags(r *http.Request) (interface{}, error) {
	ctx := r.Context()

	projectID, err := url.ProjectIDFromRequest(r)
	if err != nil {
		return nil, err
	}

	var req tagsRemovalRequest
	if err := util.UnmarshalFromReaderCloser(r.Body, &req); err != nil {
		return nil, err
	}

	if err := ep.tagManager.RemoveTags(ctx, projectID, req.IDs); err != nil {
		return nil, err
	}

	return ep.doGetTags(ctx, projectID)
}

func (ep *TagEndpoints) doGetTags(ctx context.Context, projectID projectid.ID) ([]*clienttag.Tag, error) {
	tags, err := ep.tagManager.GetTags(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return ep.tagConverter.ConvertTags(ctx, tags), nil
}

func (ep *TagEndpoints) toDomainTags(clientTags []*clienttag.Tag) []*tag.Tag {
	result := make([]*tag.Tag, 0, len(clientTags))
	for _, ct := range clientTags {
		result = append(result, &tag.Tag{
			ID:   ct.ID,
			Text: ct.Text,
		})
	}

	return result
}
