package dotx

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/emicklei/dot"
)

type compositeGraphKind int

const (
	// SameGraph means that the composite graph will be a cluster within the graph.
	SameGraph compositeGraphKind = iota
	// ExternalGraph means the the composite graph will be exported on its own, linked by the node within the graph
	ExternalGraph
)

// Composite is a graph and node to create abstractions in graphs.
type Composite struct {
	*dot.Graph
	outerNode   dot.Node
	outerGraph  *dot.Graph
	dotFilename string
	kind        compositeGraphKind
}

// NewComposite creates a Composite abstraction that is represented as a Node (box3d shape) in the graph.
// The kind determines whether the graph of the composite is embedded (same graph) or external.
func NewComposite(id string, g *dot.Graph, kind compositeGraphKind) *Composite {
	var innerGraph *dot.Graph
	if kind == SameGraph {
		innerGraph = g.Subgraph(id, dot.ClusterOption{})
	} else {
		innerGraph = dot.NewGraph(dot.Directed)
	}
	sub := &Composite{
		Graph:      innerGraph,
		outerNode:  g.Node(id).Attr("shape", "box3d"),
		outerGraph: g,
		kind:       kind,
	}
	sub.ExportName(id)
	return sub
}

// Attr sets label=value and returns the Node in the graph
func (s *Composite) Attr(label string, value interface{}) dot.Node {
	return s.outerNode.Attr(label, value)
}

// ExportName argument name will be used for the .dot export and the HREF link using svg
// So if name = "my example" then export will create "my_example.dot" and the link will be "my_example.svg"
func (s *Composite) ExportName(name string) {
	href := strings.ReplaceAll(name, " ", "_") + ".svg"
	dot := strings.ReplaceAll(name, " ", "_") + ".dot"
	s.outerNode.Attr("href", href)
	s.dotFilename = dot
}

// Input creates an edge.
// If the from Node is part of the parent graph then the edge is added to the parent graph.
// If the from Node is part of the composite then the edge is added to the inner graph.
func (s *Composite) Input(id string, from dot.Node) dot.Edge {
	if s.Graph.HasNode(from) {
		// edge on innergraph
		return s.connect(id, true, from)
	}
	// edge on outergraph
	return from.Edge(s.outerNode).Label(id)
}

// Output creates an edge.
// If the to Node is part of the parent graph then the edge is added to the parent graph.
// If the to Node is part of the composite then the edge is added to the inner graph.
func (s *Composite) Output(id string, to dot.Node) dot.Edge {
	if s.Graph.HasNode(to) {
		// edge on innergraph
		return s.connect(id, false, to)
	}
	// edge on outergraph
	return s.outerNode.Edge(to).Label(id)
}

func (s *Composite) connect(portName string, isInput bool, inner dot.Node) dot.Edge {
	port := s.Node(portName).Attr("shape", "point")
	if isInput {
		return s.EdgeWithPorts(port, inner, "s", "n").Attr("taillabel", portName)
	} else {
		// is output
		return s.EdgeWithPorts(inner, port, "s", "n").Attr("headlabel", portName)
	}
}

// ExportFile creates a DOT file using the default name (based on name) or overridden using ExportName().
func (s *Composite) ExportFile() error {
	if s.kind != ExternalGraph {
		return errors.New("ExportFile is only applicable to a ExternalGraph Composite")
	}
	return os.WriteFile(s.dotFilename, []byte(s.Graph.String()), os.ModePerm)
}

// Export writes the DOT file for a Composite after building the content (child) graph using the build function.
// Use ExportName() on the Composite to modify the filename used.
// If writing of the file fails then a warning is logged.
func (s *Composite) Export(build func(g *dot.Graph)) *Composite {
	build(s.Graph)
	if err := s.ExportFile(); err != nil {
		log.Println("WARN: dotx.Composite.Export failed", err)
	}
	return s
}
