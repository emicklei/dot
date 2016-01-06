package dot

import (
	"bytes"
	"fmt"
	"strings"
)

// AttributesMap holds attribute=value pairs.
type AttributesMap struct {
	attributes map[string]interface{}
}

// Attr sets the value for an attribute (unless empty).
func (a AttributesMap) Attr(label string, value interface{}) {
	if len(label) == 0 {
		return
	}
	a.attributes[label] = value
}

type Node struct {
	AttributesMap
	graph *Digraph
	id    string
	seq   int
}

// Attr sets label=value and return the Node
func (n Node) Attr(label string, value interface{}) Node {
	n.AttributesMap.Attr(label, value)
	return n
}

// Box sets the attribute "shape" to "box"
func (n Node) Box() Node {
	return n.Attr("shape", "box")
}

// Edge sets label=value and return the Edge
func (n Node) Edge(o Node, labels ...string) Edge {
	return n.graph.Edge(n, o, labels...)
}

// Edge represents a graph edge between two Nodes.
type Edge struct {
	AttributesMap
	graph    *Digraph
	from, to Node
}

// Attr sets label=value and return the Egde
func (e Edge) Attr(label string, value interface{}) Edge {
	e.AttributesMap.Attr(label, value)
	return e
}

type Digraph struct {
	seq       int
	nodes     map[string]Node
	edgesFrom map[string][]Edge
}

func NewDigraph() *Digraph {
	return &Digraph{
		nodes:     map[string]Node{},
		edgesFrom: map[string][]Edge{},
	}
}

// Node returns the node created with this id or creates a new node if absent.
// This method can be used as both a constructor and accessor.
func (g *Digraph) Node(id string) Node {
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

// Edge creates a new edge from n1 to n2. Nodes can be have multiple edges to the same other node (or itself).
// If one or more labels are given then the "label" attribute is set to the concatenation.
func (g *Digraph) Edge(n1, n2 Node, labels ...string) Edge {
	e := Edge{
		from:          n1,
		to:            n2,
		AttributesMap: AttributesMap{attributes: map[string]interface{}{}},
		graph:         g}
	g.edgesFrom[n1.id] = append(g.edgesFrom[n1.id], e)
	if len(labels) > 0 {
		e.Attr("label", strings.Join(labels, ","))
	}
	return e
}

// String returns the source in dot notation.
func (g Digraph) String() string {
	b := new(bytes.Buffer)
	b.WriteString("digraph {\n")
	for _, each := range g.nodes {
		fmt.Fprintf(b, "\tnode")
		if len(each.attributes) > 0 {
			b.WriteString(" [")
			first := true
			for label, value := range each.attributes {
				if !first {
					fmt.Fprintf(b, ", ")
				}
				fmt.Fprintf(b, "%s=%q", label, value)
				first = false
			}
			b.WriteString("]")
		}
		fmt.Fprintf(b, "; n%d;\n", each.seq)
	}
	for _, all := range g.edgesFrom {
		for _, each := range all {
			fmt.Fprintf(b, "\tn%d -> n%d", each.from.seq, each.to.seq)
			if len(each.attributes) > 0 {
				b.WriteString(" [")
				first := true
				for label, value := range each.attributes {
					if !first {
						fmt.Fprintf(b, ", ")
					}
					fmt.Fprintf(b, "%s=%q", label, value)
					first = false
				}
				b.WriteString("]")
			}
			b.WriteString(";\n")
		}
	}
	b.WriteString("}")
	return b.String()
}
