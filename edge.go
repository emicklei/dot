package dot

// Edge represents a graph edge between two Nodes.
type Edge struct {
	AttributesMap
	graph    *Graph
	from, to Node
}

// Attr sets key=value and returns the Egde.
func (e Edge) Attr(key string, value interface{}) Edge {
	e.AttributesMap.Attr(key, value)
	return e
}

// Label sets "label"=value and returns the Edge.
// Same as Attr("label",value)
func (e Edge) Label(value interface{}) Edge {
	e.AttributesMap.Attr("label", value)
	return e
}

// Edge returns a new Edge between the "to" node of this Edge and the argument Node.
func (e Edge) Edge(to Node, labels ...string) Edge {
	return e.graph.Edge(e.to, to, labels...)
}
