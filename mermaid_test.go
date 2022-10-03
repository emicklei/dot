package dot

import (
	"testing"
)

func TestMermaidSimple(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("e1").Label("E1")
	n2 := di.Node("e2")
	n1.Edge(n2, "what").Attr("x", "y")
	t.Log(MermaidGraph(di, MermaidTopDown))
}
