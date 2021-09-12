package apitest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/suite"
)

type TagsTestSuite struct {
	Suite
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, &TagsTestSuite{})
}

func (s *TagsTestSuite) TestCreateTags() {
	s.createTags(projectid.New(), &clienttag.Tag{
		ID:   tagid.New(),
		Text: "tag1",
	})
}

type tagsCreationRequest struct {
	Tags []*clienttag.Tag `json:"tags"`
}

func (s *TagsTestSuite) createTags(projectID projectid.ID, tags ...*clienttag.Tag) {
	b, err := json.Marshal(tagsCreationRequest{
		Tags: tags,
	})
	s.Require().NoError(err)

	resp := s.CallEndpoint(fmt.Sprintf("/projects/%s/tags/create", projectID), b)

	s.Require().EqualValues(http.StatusOK, resp.StatusCode, "%s", spew.Sdump(resp))
}
