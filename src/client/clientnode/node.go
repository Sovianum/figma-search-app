package clientnode

import "github.com/Sovianum/figma-search-app/src/figma/figmaclient"

type Node struct {
	NodeID figmaclient.NodeID `json:"nodeId"`
	User   string             `json:"user"`
}
