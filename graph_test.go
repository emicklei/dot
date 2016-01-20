package dot

import "testing"

func TestEmpty(t *testing.T) {
	di := NewDigraph()
	if got, want := di.String(), `digraph {
}
`; got != want {
		t.Fail()
	}
}

func TestEmptyWithIDAndAttributes(t *testing.T) {
	di := NewDigraph()
	di.ID("test")
	di.Attr("style", "filled")
	di.Attr("color", "lightgrey")
	if got, want := di.String(), `digraph {
	ID="test";
	style="filled";
	color="lightgrey";
}
`; got != want {
		println(got)
		t.Fail()
	}
}

func TestTwoConnectedNodes(t *testing.T) {
	di := NewDigraph()
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := di.String(), `digraph {
	node [label="A"]; n1;
	node [label="B"]; n2;
	n1 -> n2;
}
`; got != want {
		println(got)
		t.Fail()
	}
}

func TestSubgraph(t *testing.T) {
	di := NewGraph(Directed)
	sub := di.Subgraph("test")
	sub.Attr("style", "filled")
	if got, want := di.String(), `digraph {
subgraph {
	ID="s0";
	label="test";
	style="filled";
}
}
`; got != want {
		println(got)
		t.Fail()
	}
}
