package dot

import (
	"testing"
)

func TestNode_Box(t *testing.T) {
	g := NewGraph(Directed)
	n := g.Node("A")
	n.Box()
	if n.Value("shape") != "box" {
		t.Fail()
	}
}

func TestNode_Label(t *testing.T) {
	g := NewGraph(Directed)
	n := g.Node("A")
	n.Label("42")
	if n.Value("label") != "42" {
		t.Fail()
	}
}
