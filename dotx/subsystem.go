package dotx

import (
	"os"
	"strings"

	"github.com/emicklei/dot"
)

type subsystemKind int

const (
	// SameGraph means that the subsystem graph will be a cluster within the graph.
	SameGraph subsystemKind = iota
	// ExternalGraph means the the subsystem graph will be exported on its own, linked by the node within the graph
	ExternalGraph
)

// Subsystem is a graph and node to create abstractions in graphs.
type Subsystem struct {
	*dot.Graph
	outerNode   dot.Node
	outerGraph  *dot.Graph
	dotFilename string
}

// NewSubsystem creates a Subsystem abstraction that is represented as a Node (box3d shape) in the graph.
// The kind determines whether the graph of the subsystem is embedded (same graph) or external.
func NewSubsystem(id string, g *dot.Graph, kind subsystemKind) *Subsystem {
	var innerGraph *dot.Graph
	if kind == SameGraph {
		innerGraph = g.Subgraph(id, dot.ClusterOption{})
	} else {
		innerGraph = dot.NewGraph(dot.Directed)
	}
	sub := &Subsystem{
		Graph:      innerGraph,
		outerNode:  g.Node(id).Attr("shape", "box3d"),
		outerGraph: g,
	}
	sub.ExportName(id)
	return sub
}

// Attr sets label=value and returns the Node in the graph
func (s *Subsystem) Attr(label string, value interface{}) dot.Node {
	return s.outerNode.Attr(label, value)
}

// This name will be used for the .dot export and the HREF link using svg
// So if name = "my example" then export will create "my_example.dot" and the link will be "my_example.svg"
func (s *Subsystem) ExportName(name string) {
	href := strings.ReplaceAll(name, " ", "_") + ".svg"
	dot := strings.ReplaceAll(name, " ", "_") + ".dot"
	s.outerNode.Attr("href", href)
	s.dotFilename = dot
}

// Input creates an edge.
// If the from Node is part of the parent graph then the edge is added to the parent graph.
// If the from Node is part of the subsystem then the edge is added to the inner graph.
func (s *Subsystem) Input(id string, from dot.Node) dot.Edge {
	if _, ok := s.FindNodeById(from.ID()); ok {
		// edge on innergraph
		return s.connect(id, true, from)
	}
	// edge on outergraph
	return from.Edge(s.outerNode).Label(id)
}

// Output creates an edge.
// If the to Node is part of the parent graph then the edge is added to the parent graph.
// If the to Node is part of the subsystem then the edge is added to the inner graph.
func (s *Subsystem) Output(id string, to dot.Node) dot.Edge {
	if _, ok := s.FindNodeById(to.ID()); ok {
		// edge on innergraph
		return s.connect(id, false, to)
	}
	// edge on outergraph
	return s.outerNode.Edge(to).Label(id)
}

func (s *Subsystem) connect(portName string, isInput bool, inner dot.Node) dot.Edge {
	port := s.Node(portName).Attr("shape", "point")
	if isInput {
		return s.EdgeWithPorts(port, inner, "s", "n").Attr("taillabel", portName)
	} else {
		// is output
		return s.EdgeWithPorts(inner, port, "s", "n").Attr("headlabel", portName)
	}
}

// ExportFile creates a DOT file using the default name (based on name) or set using ExportName.
func (s *Subsystem) ExportFile() error {
	return os.WriteFile(s.dotFilename, []byte(s.Graph.String()), os.ModePerm)
}
