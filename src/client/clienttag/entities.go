package clienttag

import "github.com/Sovianum/figma-search-app/src/domain/tag/tagid"

type Tag struct {
	ID   tagid.ID `json:"id"`
	Text string   `json:"text"`
}
