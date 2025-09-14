package dotx

import (
	"os"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func writeDot() bool {
	return os.Getenv("WRITE_DOT_TEST_OUTPUT") != "" // set to any value to write dot files
}

func TestExampleSubsystemSameGraph(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	c1 := g.Node("component")

	sub := NewComposite("subsystem", g, SameGraph)
	sub.Input("in1", c1)
	sub.Input("in2", c1)
	sub.Output("out2", c1)

	sc1 := sub.Node("subcomponent 1")
	sc2 := sub.Node("subcomponent 2")
	sub.Input("in1", sc1)
	sub.Input("in2", sc2)
	sub.Output("out2", sc2)

	sc1.Edge(sc2)

	sub2 := NewComposite("subsystem2", sub.Graph, SameGraph)
	sub2.Input("in3", sc1)
	sub2.Output("out3", sc2)

	sub3 := sub2.Node("subcomponent 3")
	sub2.Input("in3", sub3)

	expected := `digraph  {
	subgraph cluster_s2 {
		subgraph cluster_s9 {
			label="subsystem2";
			n11[label="in3",shape="point"];
			n12[label="out3",shape="point"];
			n13[label="subcomponent 3"];
			n11:s->n13:n[taillabel="in3"];
			
		}
		label="subsystem";
		n4[label="in1",shape="point"];
		n5[label="in2",shape="point"];
		n6[label="out2",shape="point"];
		n7[label="subcomponent 1"];
		n8[label="subcomponent 2"];
		n10[href="subsystem2.svg",label="subsystem2",shape="box3d"];
		n4:s->n7:n[taillabel="in1"];
		n5:s->n8:n[taillabel="in2"];
		n7->n8;
		n7->n10[label="in3"];
		n8:s->n6:n[headlabel="out2"];
		n10->n8[label="out3"];
		
	}
	
	n1[label="component"];
	n3[href="subsystem.svg",label="subsystem",shape="box3d"];
	n1->n3[label="in1"];
	n1->n3[label="in2"];
	n3->n1[label="out2"];
	
}
`
	if got, want := g.String(), expected; got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
	if writeDot() {
		os.WriteFile("TestExampleSubsystemSameGraph.dot", []byte(g.String()), 0666)
	}
}

func TestExampleSubsystemExternalGraph(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	c1 := g.Node("component")

	sub := NewComposite("subsystem", g, ExternalGraph)
	sub.Input("in1", c1)
	sub.Input("in2", c1)
	sub.Output("out2", c1)

	sub.Export(func(g *dot.Graph) {
		sc1 := sub.Node("subcomponent 1")
		sc2 := sub.Node("subcomponent 2")
		sub.Input("in1", sc1)
		sub.Input("in2", sc2)
		sub.Output("out2", sc2)
		sc1.Edge(sc2)

		sub2 := NewComposite("subsystem2", sub.Graph, ExternalGraph)
		sub2.Export(func(g *dot.Graph) {
			sub2.Input("in3", sc1)
			sub2.Output("out3", sc2)
			sub3 := sub2.Node("subcomponent 3")
			sub2.Input("in3", sub3)
		})
	})
	expected := `digraph  {
	
	n1[label="component"];
	n2[href="subsystem.svg",label="subsystem",shape="box3d"];
	n1->n2[label="in1"];
	n1->n2[label="in2"];
	n2->n1[label="out2"];
	
}
`
	if got, want := g.String(), expected; got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
	if writeDot() {
		os.WriteFile("TestExampleSubsystemExternalGraph.dot", []byte(g.String()), 0666)
	} else {
		if err := os.Remove("subsystem.dot"); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove("subsystem2.dot"); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAttrOnSubsystem(t *testing.T) {
	s := NewComposite("test", dot.NewGraph(), SameGraph)
	s.Attr("shape", "box3d")
	if !strings.Contains(s.String(), "test") { // dont care about structure, dot has tested that
		t.Fail()
	}
}

func TestWarninOnExport(t *testing.T) {
	s := NewComposite("/////fail", dot.NewGraph(), SameGraph)
	s.Export(func(g *dot.Graph) {})
}

func TestCompositeWithUnusedIOSameGraph(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	c1 := g.Node("component")
	sub := NewComposite("subsystem", g, SameGraph)
	sub.Input("in", c1)
	sub.Output("out", c1)

	if writeDot() {
		os.WriteFile("TestCompositeWithUnusedIOSameGraph.dot", []byte(g.String()), 0666)
	}
}

func TestConnectToComposites(t *testing.T) {
	g := dot.NewGraph()
	c1 := NewComposite("c1", g, SameGraph)
	c2 := NewComposite("c2", g, SameGraph)
	e := c1.Input("in", c2)
	if e.From().ID() != c2.outerNode.ID() {
		t.Fail()
	}
	f := c1.Output("out", c2)
	if f.To().ID() != c2.outerNode.ID() {
		t.Fail()
	}
}
