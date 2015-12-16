## dot - DSL in Go for the graphviz dot language

[DOT language](http://www.graphviz.org/doc/info/lang.html)

	package main
	
	import (
		"fmt"	
		"github.com/emicklei/dot"
	)
	
	// go run main.go | dot -Tpng  > test.png && open test.png
	
	func main() {
		g := dot.NewDigraph()
		n1 := g.Node("coding")
		n2 := g.Node("testing a little")
		n2.Attr("shape", "box")
	
		g.Edge(n1, n2)
		e := g.Edge(n2, n1)
		e.Attr("color", "red")
	
		fmt.Println(g.String())
	}

Output

	digraph {
		node [label="coding",]; n1;
		node [label="testing a little",shape="box",]; n2;
		n1 -> n2 [];
		n2 -> n1 [color="red",];
	}

(c) 2015, http://ernestmicklei.com. MIT License