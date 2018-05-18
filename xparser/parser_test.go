package xparser

import (
	"strings"
	"testing"
)

func TestParseSample(t *testing.T) {
	p := NewParser(strings.NewReader(sample))
	g, err := p.ParseGraph()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", g)
}

var sample = `
digraph {
	graph [_draw_="c 9 -#fffffe00 C 7 -#ffffff P 4 0 0 0 297 1424.14 297 1424.14 0 ",
		bb="0,0,1424.1,297",
		xdotversion=1.7
	];
	n6	 [URL="",
		_draw_="c 7 -#000000 e 104.64 18 104.78 18 ",
		_ldraw_="F 14 11 -Times-Roman c 7 -#000000 T 104.64 14.3 0 145 26 --dagobert-gross-price-con- ",
		color="",
		height=0.5,
		label="-gross-price-con-",
		labeltooltip="",
		pos="104.64,18",
		shape="",
		width=2.9067];
	n5 -> n6	 [_draw_="c 7 -#000000 B 4 328.53 90.14 283.07 75.49 213.38 53.04 163.84 37.08 ",
		_hdraw_="S 5 -solid c 7 -#000000 C 7 -#000000 P 3 164.81 33.71 154.22 33.97 162.66 40.37 ",
		_ldraw_="F 14 11 -Times-Roman c 7 -#000000 T 276.64 57.8 0 40 6 -pubsub ",
		label=pubsub,
		lp="276.64,61.5",
		pos="e,154.22,33.975 328.53,90.143 283.07,75.493 213.38,53.039 163.84,37.077"];
}`
