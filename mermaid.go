package dot

import (
	"fmt"
	"strings"
)

const (
	MermaidTopToBottom = iota
	MermaidTopDown
	MermaidBottomToTop
	MermaidRightToLeft
	MermaidLeftToRight
)

func Mermaid(g *Graph, orientation int) string {
	sb := new(strings.Builder)
	sb.WriteString("graph ")
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
	mermaidEnd(sb)
	for k, v := range g.edgesFrom {
		for _, each := range v {
			sb.WriteString(k)
			if label := each.from.GetAttr("label"); label != nil {
				fmt.Fprintf(sb, "(%s)", label.(string))
			}
			//if g.graphType == Directed TODO
			if label := each.GetAttr("label"); label != nil {
				fmt.Fprintf(sb, "-->|%s|", label.(string))
			} else {
				sb.WriteString("-->")
			}
			sb.WriteString(each.to.id)
			mermaidEnd(sb)
		}
	}
	return sb.String()
}

func mermaidEnd(sb *strings.Builder) {
	sb.WriteString(";\n")
}
