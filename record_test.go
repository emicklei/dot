package dot

import (
	"fmt"
	"testing"
)

func TestSimpleRecord(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.Field("a")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="a",shape="record"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestSimpleMRecordWithFieldID(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.MRecord()
	rb.FieldWithId("a", "a1")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="<a1> a",shape="mrecord"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoColumnsRecord(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.Field("a").Field("b")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="a|b",shape="record"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoColumnsNestedRecord(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.Field("a")
	rb.Nesting(func() {
		rb.Field("b")
		rb.Field("c")
	})
	rb.Field("d")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="a|{b|c}|d",shape="record"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestStack(t *testing.T) {
	one := recordLabel{}
	two := recordLabel{}
	s := new(stack)
	s.push(one)
	s.push(two)
	if got, want := s.pop(), two; &got == &want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := s.pop(), one; &got == &want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := len(*s), 0; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

// https://graphviz.org/doc/info/shapes.html#record
/*
	digraph structs {
	    node [shape=record];
	    struct1 [label="<f0> left|<f1> mid&#92; dle|<f2> right"];
	    struct2 [label="<f0> one|<f1> two"];
	    struct3 [label="hello&#92;nworld |{ b |{c|<here> d|e}| f}| g | h"];
	    struct1:f1 -> struct2:f0;
	    struct1:f2 -> struct3:here;
	}
*/
func ExampleRecordBuilder() {
	g := NewGraph(Directed)

	r1 := NewRecordBuilder(g.Node("struct1"))
	r1.FieldWithId("left", "f0")
	r1.FieldWithId("mid dle", "f1")
	r1.FieldWithId("right", "f2")

	r2 := NewRecordBuilder(g.Node("struct2"))
	r2.FieldWithId("one", "f0")

	r3 := NewRecordBuilder(g.Node("struct3"))
	r3.Field("hello world")
	r3.Nesting(func() {
		r3.Field("b")
		r3.Nesting(func() {
			r3.Field("c")
			r3.FieldWithId("d", "here")
			r3.Field("e")
		})
		r3.Field("f")
	})
	r3.Field("g")
	r3.Field("h")

	g.Node("struct1").Edge(g.Node("struct2"))

	fmt.Println(g.String())
}
