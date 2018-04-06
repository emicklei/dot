package dot

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	Strict     = "strict"
	Undirected = "graph"
	Directed   = "digraph"
	Sub        = "subgraph"
)

type Graph struct {
	AttributesMap
	id        string
	graphType string
	seq       int
	nodes     map[string]Node
	edgesFrom map[string][]Edge
	subgraphs map[string]*Graph
}

func NewDigraph() *Graph {
	return NewGraph(Directed)
}

func NewGraph(graphType string) *Graph {
	return &Graph{
		AttributesMap: AttributesMap{attributes: map[string]interface{}{}},
		graphType:     graphType,
		nodes:         map[string]Node{},
		edgesFrom:     map[string][]Edge{},
		subgraphs:     map[string]*Graph{},
	}
}

func (g *Graph) ID(newID string) *Graph {
	g.id = newID
	return g
}

// Subgraph returns the Graph with the given label ; creates one if absent.
func (g *Graph) Subgraph(label string) *Graph {
	sub, ok := g.subgraphs[label]
	if ok {
		return sub
	}
	sub = NewGraph(Sub)
	sub.Attr("label", label)
	sub.ID(fmt.Sprintf("s%d", len(g.subgraphs)))
	g.subgraphs[label] = sub
	return sub
}

// Node returns the node created with this id or creates a new node if absent.
// This method can be used as both a constructor and accessor.
func (g *Graph) Node(id string) Node {
	n, ok := g.nodes[id]
	if ok {
		return n
	}
	// create a new
	g.seq++
	n = Node{
		id:  id,
		seq: g.seq,
		AttributesMap: AttributesMap{attributes: map[string]interface{}{
			"label": id}},
		graph: g,
	}
	g.nodes[id] = n
	return n
}

// Edge creates a new edge between two nodes.
// Nodes can be have multiple edges to the same other node (or itself).
// If one or more labels are given then the "label" attribute is set to the edge.
func (g *Graph) Edge(fromNode, toNode Node, labels ...string) Edge {
	e := Edge{
		from:          fromNode,
		to:            toNode,
		AttributesMap: AttributesMap{attributes: map[string]interface{}{}},
		graph:         g}
	g.edgesFrom[fromNode.id] = append(g.edgesFrom[fromNode.id], e)
	if len(labels) > 0 {
		e.Attr("label", strings.Join(labels, ","))
	}
	return e
}

// String returns the source in dot notation.
func (g Graph) String() string {
	b := new(bytes.Buffer)
	g.Write(b)
	return b.String()
}

// String returns the source in dot notation.
func (g Graph) Write(b io.Writer) {
	fmt.Fprintf(b, "%s{", g.graphType)
	if len(g.id) > 0 {
		fmt.Fprintf(b, "ID=%q;", g.id)
	}
	// subgraphs
	for _, each := range g.subgraphs {
		each.Write(b)
	}
	// graph attributes
	appendSortedMap(g.AttributesMap.attributes, false, b)
	// graph nodes
	for _, each := range g.nodes {
		fmt.Fprintf(b, "node")
		appendSortedMap(each.attributes, true, b)
		fmt.Fprintf(b, "n%d;", each.seq)
	}
	// graph edges
	for _, all := range g.edgesFrom {
		for _, each := range all {
			fmt.Fprintf(b, "n%d->n%d", each.from.seq, each.to.seq)
			appendSortedMap(each.attributes, true, b)
			fmt.Fprint(b, ";")
		}
	}
	fmt.Fprintf(b, "}")
}

func appendSortedMap(m map[string]interface{}, mustBracket bool, b io.Writer) {
	if len(m) == 0 {
		return
	}
	if mustBracket {
		fmt.Fprint(b, "[")
	}
	first := true
	// first collect keys
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.StringSlice(keys).Sort()

	for _, k := range keys {
		if !first {
			fmt.Fprintf(b, ",")
		}
		fmt.Fprintf(b, "%s=%q", k, m[k])
		first = false
	}
	if mustBracket {
		fmt.Fprint(b, "]")
	}
}
