package dot

import (
	"testing"
)

func TestSimpleRecord(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.AddField("a")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="a",shape="record"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoColumnsRecord(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.AddField("a")
	rb.AddField("b")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="a | b",shape="record"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoColumnsNestedRecord(t *testing.T) {
	g := NewGraph(Directed)

	rb := NewRecordBuilder(g.Node("r"))
	rb.AddField("a")
	rb.FlipWhile(func() {
		rb.AddField("b")
		rb.AddField("c")
	})
	rb.AddField("d")
	rb.Build()

	if got, want := flatten(g.String()), `digraph  {n1[label="a | b",shape="record"];}`; got != want {
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
