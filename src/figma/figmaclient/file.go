package figmaclient

import "time"

type FileID string

type File struct {
	Document     *NodeDocument             `json:"document"`
	Components   map[NodeID]*NodeComponent `json:"components"`
	Name         string                    `json:"name"`
	LastModified time.Time                 `json:"lastModified"`
	ThumbnailUrl string                    `json:"thumbnailUrl"`
	Version      string                    `json:"version"`
}
