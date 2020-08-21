package dot

import (
	"testing"
)

func TestEdgeStyleHelpers(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "solid", want: `digraph  {n1[label="A"];n2[label="B"];n1->n2[style="solid"];}`},
		{input: "bold", want: `digraph  {n1[label="A"];n2[label="B"];n1->n2[style="bold"];}`},
		{input: "dashed", want: `digraph  {n1[label="A"];n2[label="B"];n1->n2[style="dashed"];}`},
		{input: "dotted", want: `digraph  {n1[label="A"];n2[label="B"];n1->n2[style="dotted"];}`},
	}

	for _, tc := range tests {

		di := NewGraph(Directed)
		n1 := di.Node("A")
		n2 := di.Node("B")

		switch tc.input {
		case "solid":
			di.Edge(n1, n2).Solid()
		case "bold":
			di.Edge(n1, n2).Bold()
		case "dashed":
			di.Edge(n1, n2).Dashed()
		case "dotted":
			di.Edge(n1, n2).Dotted()
		}

		if got, want := flatten(di.String()), tc.want; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}
