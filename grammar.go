package dot

import (
	"bytes"
	"fmt"
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

type node struct {
	AttributesMap
	id  string
	seq int
}

func (n node) Attr(label string, value interface{}) node {
	n.AttributesMap.Attr(label, value)
	return n
}

type edge struct {
	AttributesMap
	from, to node
}

func (e edge) Attr(label string, value interface{}) edge {
	e.AttributesMap.Attr(label, value)
	return e
}

type Digraph struct {
	seq       int
	nodes     map[string]node
	edgesFrom map[string][]edge
}

func NewDigraph() *Digraph {
	return &Digraph{
		nodes:     map[string]node{},
		edgesFrom: map[string][]edge{},
	}
}

// Node creates a new with label set to id
func (g *Digraph) Node(id string) node {
	g.seq++
	return node{id: id, seq: g.seq, AttributesMap: AttributesMap{attributes: map[string]interface{}{
		"label": id,
	}}}
}

// Node creates a new with from set to n1 and to set to n2
func (g Digraph) Edge(n1, n2 node) edge {
	e := edge{from: n1, to: n2, AttributesMap: AttributesMap{attributes: map[string]interface{}{}}}
	g.nodes[n1.id] = n1
	g.nodes[n2.id] = n2
	g.edgesFrom[n1.id] = append(g.edgesFrom[n1.id], e)
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
