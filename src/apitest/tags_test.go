package apitest

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/Sovianum/figma-search-app/src/domain/files/fileid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
	"github.com/stretchr/testify/suite"
)

type TagsTestSuite struct {
	Suite
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, &TagsTestSuite{})
}

func (s *TagsTestSuite) TestCreateTags() {
	cl := s.NewClient()

	createTags(cl, fileid.New(), &clienttag.Tag{
		ID:   tagid.New(),
		Text: "tag1",
	})
}

type tagsCreationRequest struct {
	Tags []*clienttag.Tag `json:"tags"`
}

func createTags(c *Client, fileID fileid.ID, tags ...*clienttag.Tag) {
	resp := NewQueryBuilder(c).
		Pathf("/files/%s/tags/create", fileID).
		WithRequestBody(tagsCreationRequest{
			Tags: tags,
		}).
		Post()

	if resp.Status != http.StatusOK {
		panic(errors.New(fmt.Sprintf("unexpected response code %d", resp.Status)))
	}
}
