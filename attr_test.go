package dot

import "testing"

func TestAttributesMap_Attrs(t *testing.T) {
	g := NewGraph()
	g.Attrs("l", "v", "l2", "v2")
	if got, want := g.attributes["l"], "v"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := g.attributes["l2"], "v2"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
