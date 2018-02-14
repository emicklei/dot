package dot

// Node represents a dot Node.
type Node struct {
	AttributesMap
	graph *Graph
	id    string
	seq   int
}

// Attr sets label=value and return the Node
func (n Node) Attr(label string, value interface{}) Node {
	n.AttributesMap.Attr(label, value)
	return n
}

// Box sets the attribute "shape" to "box"
func (n Node) Box() Node {
	return n.Attr("shape", "box")
}

// Edge sets label=value and returns the Edge for chaining.
func (n Node) Edge(toNode Node, labels ...string) Edge {
	return n.graph.Edge(n, toNode, labels...)
}
