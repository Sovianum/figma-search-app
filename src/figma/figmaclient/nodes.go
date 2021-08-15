package figmaclient

type NodeID string

type NodeType string

const (
	NodeTypeDocument     NodeType = "DOCUMENT"
	NodeTypeFrame        NodeType = "FRAME"
	NodeTypeGroup        NodeType = "GROUP"
	NodeTypeText         NodeType = "TEXT"
	NodeTypeComponent    NodeType = "COMPONENT"
	NodeTypeComponentSet NodeType = "COMPONENT_SET"
	NodeTypeNone         NodeType = ""
)

type Node struct {
	// one of

	Document     *NodeDocument     `json:"document,omitempty"`
	Frame        *NodeFrame        `json:"frame,omitempty"`
	Group        *NodeGroup        `json:"group,omitempty"`
	Text         *NodeText         `json:"text,omitempty"`
	Component    *NodeComponent    `json:"component,omitempty"`
	ComponentSet *NodeComponentSet `json:"componentSet,omitempty"`
	Unhandled    *NodeUnhandled    `json:"unhandled,omitempty"`
}

type NodeCommon struct {
	ID   NodeID   `json:"id"`
	Type NodeType `json:"type"`
	Name string   `json:"name"`
}

type NodeDocument struct {
	NodeCommon
	Children []*Node `json:"children,omitempty"`
}

type NodeFrame struct {
	NodeCommon
	Children []*Node `json:"children,omitempty"`
}

type NodeGroup struct {
	NodeCommon
	Children []*Node `json:"children,omitempty"`
}

type NodeText struct {
	NodeCommon
	Characters string `json:"characters"`
}

type NodeComponent struct {
	NodeCommon
	Children    []*Node `json:"children,omitempty"`
	Description string  `json:"description,omitempty"`
}

type NodeComponentSet struct {
	NodeCommon
	Children []*Node `json:"children,omitempty"`
}

type NodeInstance struct {
	NodeCommon
	Children []*Node `json:"children,omitempty"`
}

type NodeUnhandled struct {
	NodeCommon
	Children []*Node `json:"children,omitempty"`
}
