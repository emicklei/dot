package dot

// Edge represents a graph edge between two Nodes.
type Edge struct {
	AttributesMap
	graph    *Graph
	from, to Node
}

// Attr sets label=value and return the Egde
func (e Edge) Attr(label string, value interface{}) Edge {
	e.AttributesMap.Attr(label, value)
	return e
}

// Edge returns a new Edge between the "to" node of this Edge and the argument Node.
func (e Edge) Edge(o Node, labels ...string) Edge {
	return e.graph.Edge(e.to, o, labels...)
}
