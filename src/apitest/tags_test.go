package apitest

import (
	"encoding/json"
	"fmt"
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
	s.createTags(fileid.New(), &clienttag.Tag{
		ID:   tagid.New(),
		Text: "tag1",
	})
}

type tagsCreationRequest struct {
	Tags []*clienttag.Tag `json:"tags"`
}

func (s *TagsTestSuite) createTags(fileID fileid.ID, tags ...*clienttag.Tag) {
	b, err := json.Marshal(tagsCreationRequest{
		Tags: tags,
	})
	s.Require().NoError(err)

	result, err := s.handler(s.Ctx, b)
	s.Require().NoError(err)
	fmt.Println(result)
}
