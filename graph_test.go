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

import "testing"

var node *Node

func TestNodeCreating(t *testing.T) {
	node = NewNode("testNode", make(map[string]interface{}))
	if node.id != "testNode" {
		t.Log("Uncorrect node identifier")
		t.Fatal()
	}
}

func TestNodeAssign(t *testing.T) {
	childNode := NewNode("childNode", make(map[string]interface{}))
	node.AddEdge(childNode)
	if node.edges[0] != childNode {
		t.Log("Not assign edges")
		t.Fatal()
	}
}

func TestGetConnections(t *testing.T) {
	nodeList := node.GetConnections()
	if len(nodeList) < 1 {
		t.Log("Not return list of edges")
		t.Fatal()
	}
}

func TestGraphCreating(t *testing.T) {
	graph := New()
	if graph == nil {
		t.Log("Should create new graph instance")
		t.Fatal()
	}
}

func TestAddingNode(t *testing.T) {
	graph := New()
	graph.AddNode("testNode", make(map[string]interface{}))
	if graph.nodes["testNode"] == nil {
		t.Log("Not add node to graph")
		t.Fatal()
		if graph.nodes["testNode"].id != "testNode" {
			t.Log("Bad node")
			t.Fatal()
		}
	}
}

func TestAttachingNode(t *testing.T) {
	graph := New()
	testNode := NewNode("testNode", make(map[string]interface{}))
	graph.AttachNode(testNode)
	if graph.nodes["testNode"] != testNode {
		t.Log("Not correctly attaching node to graph")
		t.Fatal()
	}
}
