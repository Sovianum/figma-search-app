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
	projectID := projectid.New()
	tagID := tagid.New()

	requestTag := &clienttag.Tag{
		ID:   tagID,
		Text: "tag1",
	}

	s.createTags(projectID, requestTag)

	tags := s.getTags(projectID)
	s.Require().Len(tags, 1)

	s.EqualValues(tagID, tags[0].ID)
	s.EqualValues("tag1", tags[0].Text)
}

func (s *TagsTestSuite) getTags(projectId projectid.ID) []*clienttag.Tag {
	resp := s.CallEndpoint(fmt.Sprintf("/projects/%s/tags/get", projectId), []byte("{}"))
	s.Require().EqualValues(http.StatusOK, resp.StatusCode, "%s", spew.Sdump(resp))

	var result []*clienttag.Tag
	s.Require().NoError(json.Unmarshal([]byte(resp.Body), &result))

	return result
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
