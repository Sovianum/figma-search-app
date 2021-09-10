package tag

import (
	"github.com/Sovianum/figma-search-app/src/domain/nodes/nodeid"
	"github.com/Sovianum/figma-search-app/src/domain/project/projectid"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagid"
)

type Tag struct {
	ID   tagid.ID
	Text string
}

type TagNodesAction string

const (
	TagNodesActionTag   TagNodesAction = "tag"
	TagNodesActionUntag TagNodesAction = "untag"
)

type TagNodesQuery struct {
	Action    TagNodesAction
	ProjectID projectid.ID
	NodeIDs   []nodeid.ID
	TagIDs    []tagid.ID
}
