package figmaclient

import (
	"encoding/json"
	"errors"
)

type typedNodeProjection struct {
	Type NodeType `json:"type"`
}

func (n *Node) UnmarshalJSON(b []byte) error {
	var rawMessage json.RawMessage
	if err := json.Unmarshal(b, &rawMessage); err != nil {
		return err
	}

	var nodeProjection typedNodeProjection
	if err := json.Unmarshal(rawMessage, &nodeProjection); err != nil {
		return err
	}

	switch nodeProjection.Type {
	case NodeTypeDocument:
		doc := &NodeDocument{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.Document = doc
		return nil

	case NodeTypeFrame:
		doc := &NodeFrame{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.Frame = doc
		return nil

	case NodeTypeGroup:
		doc := &NodeGroup{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.Group = doc
		return nil

	case NodeTypeText:
		doc := &NodeText{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.Text = doc
		return nil

	case NodeTypeComponent:
		doc := &NodeComponent{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.Component = doc
		return nil

	case NodeTypeComponentSet:
		doc := &NodeComponentSet{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.ComponentSet = doc
		return nil

	case NodeTypeNone:
		panic(errors.New("type must not be empty"))

	default:
		doc := &NodeUnhandled{}
		if err := json.Unmarshal(rawMessage, doc); err != nil {
			return err
		}
		n.Unhandled = doc
		return nil
	}
}
