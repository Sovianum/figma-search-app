package api

import (
	"context"
	"net/http"

	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/Sovianum/figma-search-app/src/domain/files/fileid"
	"github.com/Sovianum/figma-search-app/src/domain/tag"
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
	panic("aaa")
}

type tagCreationRequest struct {
	Tags []*clienttag.Tag `json:"tags"`
}

func (ep *TagEndpoints) CreateTags(r *http.Request) (interface{}, error) {
	ctx := r.Context() // TODO pass custom context

	fileID, err := url.FileIDFromRequest(r)
	if err != nil {
		return nil, err
	}

	var req tagCreationRequest
	if err := util.UnmarshalFromReaderCloser(r.Body, &req); err != nil {
		return nil, err
	}

	if err := ep.tagManager.CreateTags(ctx, fileID, ep.toDomainTags(req.Tags)); err != nil {
		return nil, err
	}

	return ep.doGetTags(ctx, fileID)
}

func (ep *TagEndpoints) RemoveTags(r *http.Request) (interface{}, error) {
	panic("aaa")
}

func (ep *TagEndpoints) doGetTags(ctx context.Context, fileID fileid.ID) ([]*clienttag.Tag, error) {
	tags, err := ep.tagManager.GetTags(ctx, fileID)
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
