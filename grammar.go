package dot

import (
	"bytes"
	"fmt"
)

type attributeList struct {
	attributes map[string]interface{}
}

func (a attributeList) Attr(label string, value interface{}) {
	a.attributes[label] = value
}

type node struct {
	attributeList
	id  string
	seq int
}

type edge struct {
	attributeList
	from, to node
}

type Digraph struct {
	seq       int
	nodes     map[string]node
	edgesFrom map[string]edge
}

func NewDigraph() *Digraph {
	return &Digraph{
		nodes:     map[string]node{},
		edgesFrom: map[string]edge{},
	}
}

func (g *Digraph) Node(id string) node {
	g.seq++
	return node{id: id, seq: g.seq, attributeList: attributeList{attributes: map[string]interface{}{
		"label": id,
	}}}
}

func (g Digraph) Edge(n1, n2 node) edge {
	e := edge{from: n1, to: n2, attributeList: attributeList{attributes: map[string]interface{}{}}}
	g.nodes[n1.id] = n1
	g.nodes[n2.id] = n2
	g.edgesFrom[n1.id] = e
	return e
}

func (g Digraph) String() string {
	b := new(bytes.Buffer)
	b.WriteString("digraph {\n")
	for _, each := range g.nodes {
		fmt.Fprintf(b, "\tnode [")
		for label, value := range each.attributes {
			fmt.Fprintf(b, "%s=%q,", label, value)
		}
		fmt.Fprintf(b, "]; n%d;\n", each.seq)
	}
	for _, each := range g.edgesFrom {
		fmt.Fprintf(b, "\tn%d -> n%d [", each.from.seq, each.to.seq)
		for label, value := range each.attributes {
			fmt.Fprintf(b, "%s=%q,", label, value)
		}
		b.WriteString("];\n")
	}
	b.WriteString("}")
	return b.String()
}
