package dot

import (
	"fmt"
	"html"
	"strings"
)

const (
	MermaidTopToBottom = iota
	MermaidTopDown
	MermaidBottomToTop
	MermaidRightToLeft
	MermaidLeftToRight
)

var (
	MermaidShapeRound            = shape{"(", ")"}
	MermaidShapeStadium          = shape{"([", "])"}
	MermaidShapeSubroutine       = shape{"[[", "]]"}
	MermaidShapeCylinder         = shape{"[(", ")]"}
	MermaidShapeCirle            = shape{"((", "))"} // Deprecated: use MermaidShapeCircle instead
	MermaidShapeCircle           = shape{"((", "))"}
	MermaidShapeAsymmetric       = shape{">", "]"}
	MermaidShapeRhombus          = shape{"{", "}"}
	MermaidShapeTrapezoid        = shape{"[/", "\\]"}
	MermaidShapeTrapezoidAlt     = shape{"[\\", "/]"}
	MermaidShapeHexagon          = shape{"{{", "}}"}
	MermaidShapeParallelogram    = shape{"[/", "/]"}
	MermaidShapeParallelogramAlt = shape{"[\\", "\\]"}
	// TODO more shapes see https://mermaid.js.org/syntax/flowchart.html#node-shapes
)

type shape struct {
	open, close string
}

func MermaidGraph(g *Graph, orientation int) string {
	return diagram(g, "graph", orientation)
}

func MermaidFlowchart(g *Graph, orientation int) string {
	return diagram(g, "flowchart", orientation)
}

func escape(value string) string {
	return fmt.Sprintf(`"%s"`, html.EscapeString(value))
}

func diagram(g *Graph, diagramType string, orientation int) string {
	sb := new(strings.Builder)
	sb.WriteString(diagramType)
	sb.WriteRune(' ')
	switch orientation {
	case MermaidTopDown, MermaidTopToBottom:
		sb.WriteString("TD")
	case MermaidBottomToTop:
		sb.WriteString("BT")
	case MermaidRightToLeft:
		sb.WriteString("RL")
	case MermaidLeftToRight:
		sb.WriteString("LR")
	default:
		sb.WriteString("TD")
	}
	writeEnd(sb)
	diagramGraph(g, sb)
	for _, id := range g.sortedSubgraphsKeys() {
		each := g.subgraphs[id]
		fmt.Fprintf(sb, "subgraph %s [%s];\n", id, each.attributes["label"])
		diagramGraph(each, sb)
		fmt.Fprintln(sb, "end;")
	}
	return sb.String()
}

func diagramGraph(g *Graph, sb *strings.Builder) {
	// graph nodes
	for _, key := range g.sortedNodesKeys() {
		nodeShape := MermaidShapeRound
		each := g.nodes[key]
		if s := each.GetAttr("shape"); s != nil {
			// could be a shape or a string
			shapeString, ok := s.(string)
			if ok {
				// see if we can map the string to a shape
				mermaidShape, ok := lookupShape(shapeString)
				if ok {
					nodeShape = mermaidShape
				}
			}
			// could be a shape
			mermaidShape, ok := s.(shape)
			if ok {
				nodeShape = mermaidShape
			}
		}
		txt := "?"
		if label := each.GetAttr("label"); label != nil {
			// take string only
			slabel, ok := label.(string)
			if ok {
				txt = slabel
			}
		}
		fmt.Fprintf(sb, "\tn%d%s%s%s;\n", each.seq, nodeShape.open, escape(txt), nodeShape.close)
		if style := each.GetAttr("style"); style != nil {
			fmt.Fprintf(sb, "\tstyle n%d %s\n", each.seq, style.(string))
		}
	}
	// all edges
	// graph edges
	denoteEdge := "-->"
	if g.graphType == "graph" {
		denoteEdge = "---"
	}
	for _, each := range g.sortedEdgesFromKeys() {
		all := g.edgesFrom[each]
		for _, each := range all {
			// The edge can override the link style
			link := denoteEdge
			if l := each.GetAttr("link"); l != nil {
				// take string only
				slink, ok := l.(string)
				if ok {
					link = slink
				}
			}
			if label := each.GetAttr("label"); label != nil {
				slabel, ok := label.(string)
				if !ok {
					// make it a string
					slabel = fmt.Sprintf("%v", label)
				}
				if label != "" {
					fmt.Fprintf(sb, "\tn%d%s|%s|n%d;\n", each.from.seq, link, escape(slabel), each.to.seq)
					continue
				}
			}
			// no label
			fmt.Fprintf(sb, "\tn%d%sn%d;\n", each.from.seq, link, each.to.seq)
		}
	}
}

func writeEnd(sb *strings.Builder) {
	sb.WriteString(";\n")
}

func lookupShape(shapeName string) (shape, bool) {
	switch shapeName {
	case "round", "box":
		return MermaidShapeRound, true
	case "asymmetric":
		return MermaidShapeAsymmetric, true
	case "circle":
		return MermaidShapeCircle, true
	case "cylinder":
		return MermaidShapeCylinder, true
	case "rhombux":
		return MermaidShapeRhombus, true
	case "stadium":
		return MermaidShapeStadium, true
	case "subroutine":
		return MermaidShapeSubroutine, true
	case "trapezoid":
		return MermaidShapeTrapezoid, true
	case "trapezoid-alt":
		return MermaidShapeTrapezoidAlt, true
	case "hexagon":
		return MermaidShapeHexagon, true
	case "parallelogram":
		return MermaidShapeParallelogram, true
	case "parallelogram-alt":
		return MermaidShapeParallelogramAlt, true
	}
	return shape{}, false
}
