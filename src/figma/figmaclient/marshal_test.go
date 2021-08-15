package figmaclient_test

import (
	"encoding/json"
	"testing"

	"github.com/Sovianum/figma-search-app/src/figma/figmaclient"
	"github.com/stretchr/testify/require"
)

func TestInlineUnmarshalling(t *testing.T) {
	msg := `{"type": "DOCUMENT", "name": "doc"}`
	var node figmaclient.NodeDocument
	require.NoError(t, json.Unmarshal([]byte(msg), &node))

	require.EqualValues(t, "doc", node.Name)
}

func TestUnmarshalling(t *testing.T) {
	testCases := []struct {
		Name         string
		Msg          string
		ExpectedNode *figmaclient.Node
	}{
		{
			Name: "Document",
			Msg:  `{"type": "DOCUMENT", "name": "doc"}`,
			ExpectedNode: &figmaclient.Node{
				Document: &figmaclient.NodeDocument{
					NodeCommon: figmaclient.NodeCommon{
						Name: "doc",
					},
				},
			},
		},
		{
			Name: "Frame",
			Msg:  `{"type": "FRAME", "name": "frame"}`,
			ExpectedNode: &figmaclient.Node{
				Frame: &figmaclient.NodeFrame{
					NodeCommon: figmaclient.NodeCommon{
						Name: "frame",
					},
				},
			},
		},
		{
			Name: "Group",
			Msg:  `{"type": "GROUP", "name": "group"}`,
			ExpectedNode: &figmaclient.Node{
				Group: &figmaclient.NodeGroup{
					NodeCommon: figmaclient.NodeCommon{
						Name: "group",
					},
				},
			},
		},
		{
			Name: "Text",
			Msg:  `{"type": "TEXT", "name": "text", "characters": "chars"}`,
			ExpectedNode: &figmaclient.Node{
				Text: &figmaclient.NodeText{
					NodeCommon: figmaclient.NodeCommon{
						Name: "text",
					},
					Characters: "chars",
				},
			},
		},
		{
			Name: "Component",
			Msg:  `{"type": "COMPONENT", "name": "component", "description": "desc"}`,
			ExpectedNode: &figmaclient.Node{
				Component: &figmaclient.NodeComponent{
					NodeCommon: figmaclient.NodeCommon{
						Name: "component",
					},
					Description: "desc",
				},
			},
		},
		{
			Name: "ComponentSet",
			Msg:  `{"type": "COMPONENT_SET", "name": "component_set"}`,
			ExpectedNode: &figmaclient.Node{
				ComponentSet: &figmaclient.NodeComponentSet{
					NodeCommon: figmaclient.NodeCommon{
						Name: "component_set",
					},
				},
			},
		},
		{
			Name: "Unhandled",
			Msg:  `{"type": "TYPE", "name": "unhandled"}`,
			ExpectedNode: &figmaclient.Node{
				Unhandled: &figmaclient.NodeUnhandled{
					NodeCommon: figmaclient.NodeCommon{
						Name: "unhandled",
					},
				},
			},
		},
		{
			Name: "Nested",
			Msg:  `{"type": "DOCUMENT", "name": "doc", "children": [{"type": "FRAME", "name": "frame"}]}`,
			ExpectedNode: &figmaclient.Node{
				Document: &figmaclient.NodeDocument{
					NodeCommon: figmaclient.NodeCommon{
						Name: "doc",
					},
					Children: []*figmaclient.Node{
						{
							Frame: &figmaclient.NodeFrame{
								NodeCommon: figmaclient.NodeCommon{
									Name: "frame",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var node figmaclient.Node
			require.NoError(t, node.UnmarshalJSON([]byte(tc.Msg)))

			expectedSerialization, err := json.Marshal(tc.ExpectedNode)
			require.NoError(t, err)

			gotSerialization, err := json.Marshal(node)

			require.EqualValues(t, string(expectedSerialization), string(gotSerialization))
		})
	}
}
