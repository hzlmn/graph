// Copyright 2017 Oleh Kuchuk. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

// Node represents single graph node
type Node struct {
	params map[string]interface{}
	id     string
	edges  []*Node
}

// NewNode constructs Node structure
func NewNode(id string, params map[string]interface{}) *Node {
	return &Node{
		id:     id,
		params: params,
	}
}

// GetParams returns node params
func (n *Node) GetParams() map[string]interface{} {
	return n.params
}

// AddEdge adds connection to another Node
func (n *Node) AddEdge(node *Node) {
	for _, item := range n.edges {
		if item == node {
			return
		}
	}

	n.edges = append(n.edges, node)
}

// GetConnections return all node connections
func (n *Node) GetConnections() []*Node {
	return n.edges
}

// GetID returns current node id
func (n *Node) GetID() string {
	return n.id
}

// Graph defines graph structure
type Graph struct {
	nodes map[string]*Node
}

// New create new instance of graph structure
func New() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}

// AddNode create and add new node to graph
func (g *Graph) AddNode(nodeKey string, nodeParams map[string]interface{}) {
	for key := range g.nodes {
		if key == nodeKey {
			return
		}
	}

	g.nodes[nodeKey] = NewNode(nodeKey, nodeParams)
}

// AttachNode add existed node with it dependencies to graph
func (g *Graph) AttachNode(node *Node) {
	for key := range g.nodes {
		if key == node.GetID() {
			return
		}
	}

	g.nodes[node.GetID()] = node

	for _, childNode := range node.GetConnections() {
		g.AttachNode(childNode)
	}
}

// GetNodes allow to get list of registered nodes in graph
func (g *Graph) GetNodes() (result []*Node) {
	for _, node := range g.nodes {
		result = append(result, node)
	}
	return
}

// AddEdge adds links between two nodes by id
func (g *Graph) AddEdge(startID string, endID string) {
	g.AddNode(startID, make(map[string]interface{}))
	g.AddNode(endID, make(map[string]interface{}))
	g.nodes[startID].AddEdge(g.nodes[endID])
}

// GetNode find node by id
func (g *Graph) GetNode(nodeID string) *Node {
	return g.nodes[nodeID]
}
