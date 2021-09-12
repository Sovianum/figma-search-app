package apitest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/suite"
)

type TagsTestSuite struct {
	Suite
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, &TagsTestSuite{})
}

type dbTag struct {
	ID        tagid.ID     `dynamodbav:"id"`
	Text      string       `dynamodbav:"text"`
	ProjectID projectid.ID `dynamodbav:"pId"`
}

func (s *TagsTestSuite) TestCreateTags() {
	projectID := projectid.New()

	requestTag := &clienttag.Tag{
		ID:   tagid.New(),
		Text: "tag1",
	}

	itemMap, err := dynamodbattribute.MarshalMap(dbTag{
		ID:        requestTag.ID,
		Text:      requestTag.Text,
		ProjectID: projectID,
	})
	s.Require().NoError(err)

	s.DBMock.
		ExpectBatchWriteItem().
		WithRequest(map[string][]*dynamodb.WriteRequest{
			"tags": {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: itemMap,
					},
				},
			},
		})

	s.createTags(projectID, requestTag)
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
