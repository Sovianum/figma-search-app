package tag

import (
	"github.com/Sovianum/figma-search-app/src/domain/files/fileid"
	"github.com/Sovianum/figma-search-app/src/domain/nodes/nodeid"
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
	Action  TagNodesAction
	FileID  fileid.ID
	NodeIDs []nodeid.ID
	TagIDs  []tagid.ID
}
