package dot

import (
	"testing"
)

func TestMermaidSimple(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("e1").Label("E1")
	n2 := di.Node("e2").Attr("shape", MermaidShapeRound).Attr("style", "fill:#90EE90")
	n1.Edge(n2, "what").Attr("x", "y")
	out := flatten(MermaidGraph(di, MermaidTopDown))
	if got, want := out, `graph TD;n1("E1");n2("e2");style n2 fill:#90EE90n1-->|"what"|n2;`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestEmptyFlow(t *testing.T) {
	di := NewGraph(Directed)
	s := MermaidFlowchart(di, MermaidTopDown)
	if got, want := s, "flowchart TD;\n"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
func TestEmptyGraphLR(t *testing.T) {
	di := NewGraph(Directed)
	s := MermaidGraph(di, MermaidLeftToRight)
	if got, want := s, "graph LR;\n"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	s = MermaidGraph(di, MermaidRightToLeft)
	if got, want := s, "graph RL;\n"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	s = MermaidGraph(di, MermaidBottomToTop)
	if got, want := s, "graph BT;\n"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	s = MermaidGraph(di, 42)
	if got, want := s, "graph TD;\n"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
func TestMermaidShapes(t *testing.T) {
	di := NewGraph(Directed)
	di.Node("round").Attr("shape", MermaidShapeRound)
	di.Node("asym").Attr("shape", MermaidShapeAsymmetric)
	di.Node("circ").Attr("shape", MermaidShapeCircle)
	di.Node("cyl").Attr("shape", MermaidShapeCylinder)
	di.Node("rhom").Attr("shape", MermaidShapeRhombus)
	di.Node("stad").Attr("shape", MermaidShapeStadium)
	di.Node("sub").Attr("shape", MermaidShapeSubroutine)
	di.Node("trap").Attr("shape", MermaidShapeTrapezoid)
	di.Node("trapalt").Attr("shape", MermaidShapeTrapezoidAlt)
	s := MermaidGraph(di, MermaidLeftToRight)
	// t.Log(s)
	if got, want := flatten(s), `graph LR;n2>"asym"];n3(("circ"));n4[("cyl")];n5{"rhom"};n1("round");n6(["stad"]);n7[["sub"]];n8[/"trap"\];n9[\"trapalt"/];`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

// Deprecated: Use MermaidShapeCircle instead of MermaidShapeCirle
func TestMermaidShapeCirle(t *testing.T) {
	di := NewGraph(Directed)
	di.Node("circ").Attr("shape", MermaidShapeCirle)
	s := MermaidGraph(di, MermaidLeftToRight)
	// t.Log(s)
	if got, want := flatten(s), `graph LR;n1(("circ"));`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestUndirectedMermaid(t *testing.T) {
	un := NewGraph(Undirected)
	un.Node("love").Edge(un.Node("happinez"))
	s := MermaidFlowchart(un, MermaidLeftToRight)
	//t.Log(s)
	if got, want := flatten(s), `flowchart LR;n2("happinez");n1("love");n1---n2;`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

// example from https://mermaid.js.org/syntax/flowchart.html
// note that c1 and a2 are nodes created in their subgraphs to make the diagram match with the example.
func TestMermaidSubgraph(t *testing.T) {
	di := NewGraph(Directed)
	sub1 := di.Subgraph("one")
	sub1.Node("a1").Edge(sub1.Node("a2"))
	sub2 := di.Subgraph("two")
	sub2.Node("b1").Edge(sub2.Node("b2"))
	sub3 := di.Subgraph("THREE").Label("three")
	sub3.Node("c1").Edge(sub3.Node("c2"))

	sub3.Node("c1").Edge(sub1.Node("a2"))
	mf := MermaidFlowchart(di, MermaidLeftToRight)
	if got, want := flatten(mf), `flowchart LR;n8-->n3;subgraph THREE [three];n8("c1");n9("c2");n8-->n9;end;subgraph one [one];n2("a1");n3("a2");n2-->n3;end;subgraph two [two];n5("b1");n6("b2");n5-->n6;end;`; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestMermaidFromBoxShape(t *testing.T) {
	graph := NewGraph(Directed)
	graph.Node("A").Box()
	graph.Edge(graph.Node("A"), graph.Node("B"))

	if got, want := flatten(MermaidGraph(graph, MermaidTopDown)), `graph TD;n1("A");n2("B");n1-->n2;`; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestLookupShape(t *testing.T) {
	tests := []struct {
		name      string
		shapeName string
		wantShape shape
		wantOk    bool
	}{
		{"round", "round", MermaidShapeRound, true},
		{"box", "box", MermaidShapeRound, true},
		{"asymmetric", "asymmetric", MermaidShapeAsymmetric, true},
		{"circle", "circle", MermaidShapeCircle, true},
		{"cylinder", "cylinder", MermaidShapeCylinder, true},
		{"rhombux", "rhombux", MermaidShapeRhombus, true},
		{"stadium", "stadium", MermaidShapeStadium, true},
		{"subroutine", "subroutine", MermaidShapeSubroutine, true},
		{"trapezoid", "trapezoid", MermaidShapeTrapezoid, true},
		{"trapezoid-alt", "trapezoid-alt", MermaidShapeTrapezoidAlt, true},
		{"hexagon", "hexagon", MermaidShapeHexagon, true},
		{"parallelogram", "parallelogram", MermaidShapeParallelogram, true},
		{"parallelogram-alt", "parallelogram-alt", MermaidShapeParallelogramAlt, true},
		{"unknown", "unknown", shape{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShape, gotOk := lookupShape(tt.shapeName)
			if gotShape != tt.wantShape || gotOk != tt.wantOk {
				t.Errorf("lookupShape(%q) = (%v, %v), want (%v, %v)", tt.shapeName, gotShape, gotOk, tt.wantShape, tt.wantOk)
			}
		})
	}
}
