package xparser

import (
	"io"

	"github.com/emicklei/dot"
)

type Parser struct {
	graph *dot.Graph
}

func (p *Parser) ParseGraph() (*dot.Graph, error) {
	return new(dot.Graph), nil
}

func NewParser(r io.Reader) *Parser {
	return new(Parser)
}
