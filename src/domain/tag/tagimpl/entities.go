package tagimpl

import (
	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
)

type Tag struct {
	ID        tagid.ID     `dynamodbav:"id"`
	Text      string       `dynamodbav:"text"`
	ProjectID projectid.ID `dynamodbav:"pId"`
}
