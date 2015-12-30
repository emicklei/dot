package dot

import "testing"

func TestEmpty(t *testing.T) {
	di := NewDigraph()
	if got, want := di.String(), `digraph {
}`; got != want {
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
}`; got != want {
		println(got)
		t.Fail()
	}
}
