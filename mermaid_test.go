package dot

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestMermaidSimple(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	n1 := di.Node("e1").Label("E1")
	n2 := di.Node("e2")
	n1.Edge(n2, "what").Attr("x", "y")
	t.Log(Graph(di, TopDown))
}
