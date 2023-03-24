package dotx

import (
	"os"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestExampleSubsystemSameGraph(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	c1 := g.Node("component")

	sub := NewSubsystem("subsystem", g, SameGraph)
	sub.Input("in1", c1)
	sub.Input("in2", c1)
	sub.Output("out2", c1)

	sc1 := sub.Node("subcomponent 1")
	sc2 := sub.Node("subcomponent 2")
	sub.Input("in1", sc1)
	sub.Input("in2", sc2)
	sub.Output("out2", sc2)

	sc1.Edge(sc2)

	sub2 := NewSubsystem("subsystem2", sub.Graph, SameGraph)
	sub2.Input("in3", sc1)
	sub2.Output("out3", sc2)

	sub3 := sub2.Node("subcomponent 3")
	sub2.Input("in3", sub3)

	os.WriteFile("TestExampleSubsystemSameGraph.dot", []byte(g.String()), os.ModePerm)
}

func TestExampleSubsystemExternalGraph(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	c1 := g.Node("component")

	sub := NewSubsystem("subsystem", g, ExternalGraph)
	sub.Input("in1", c1)
	sub.Input("in2", c1)
	sub.Output("out2", c1)

	sc1 := sub.Node("subcomponent 1")
	sc2 := sub.Node("subcomponent 2")
	sub.Input("in1", sc1)
	sub.Input("in2", sc2)
	sub.Output("out2", sc2)

	sc1.Edge(sc2)

	sub2 := NewSubsystem("subsystem2", sub.Graph, ExternalGraph)
	sub2.Input("in3", sc1)
	sub2.Output("out3", sc2)

	sub3 := sub2.Node("subcomponent 3")
	sub2.Input("in3", sub3)

	// because External, we need to export both subsystems
	sub.ExportFile()
	sub2.ExportFile()

	os.WriteFile("TestExampleSubsystemExternalGraph.dot", []byte(g.String()), os.ModePerm)
}

func TestAttrOnSubsystem(t *testing.T) {
	s := NewSubsystem("test", dot.NewGraph(), SameGraph)
	s.Attr("shape", "box3d")
	if !strings.Contains(s.String(), "test") { // dont care about structure, dot has tested that
		t.Fail()
	}
}
