package dot

import "testing"

func TestEmpty(t *testing.T) {
	di := NewDigraph()
	if got, want := di.String(), `digraph{}`; got != want {
		t.Log(got)
		t.Fail()
	}
}

func TestEmptyWithIDAndAttributes(t *testing.T) {
	di := NewDigraph()
	di.ID("test")
	di.Attr("style", "filled")
	di.Attr("color", "lightgrey")
	if got, want := di.String(), `digraph{ID="test";color="lightgrey",style="filled"}`; got != want {
		t.Log(got)
		t.Fail()
	}
}

func TestTwoConnectedNodes(t *testing.T) {
	di := NewDigraph()
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := di.String(), `digraph{node[label="A"]n1;node[label="B"]n2;n1->n2;}`; got != want {
		t.Log(got)
		t.Fail()
	}
}

func TestSubgraph(t *testing.T) {
	di := NewGraph(Directed)
	sub := di.Subgraph("test")
	sub.Attr("style", "filled")
	if got, want := di.String(), `digraph{subgraph{ID="s0";label="test",style="filled"}}`; got != want {
		t.Log(got)
		t.Fail()
	}
}

func TestEdgeLabel(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("n1")
	n2 := di.Node("n2")
	n1.Edge(n2, "wat")
	if got, want := di.String(), `digraph{node[label="n1"]n1;node[label="n2"]n2;n1->n2[label="wat"];}`; got != want {
		t.Log(got)
		t.Fail()
	}
}
