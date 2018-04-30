package dot

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	di := NewGraph(Directed)
	if got, want := flatten(di.String()), `digraph  {}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEmptyWithIDAndAttributes(t *testing.T) {
	di := NewGraph(Directed)
	di.ID("test")
	di.Attr("style", "filled")
	di.Attr("color", "lightgrey")
	if got, want := flatten(di.String()), `digraph test {ID = "test";color="lightgrey";style="filled";}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTwoConnectedNodes(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("A")
	n2 := di.Node("B")
	di.Edge(n1, n2)
	if got, want := flatten(di.String()), `digraph  {node[label="A"] n1;node[label="B"] n2;n1->n2;}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestSubgraph(t *testing.T) {
	di := NewGraph(Directed)
	sub := di.Subgraph("test")
	sub.Attr("style", "filled")
	if got, want := flatten(di.String()), `digraph  {subgraph s0 {ID = "s0";label="test";style="filled";}}`; got != want {
		t.Errorf("got\n[%v] want\n[%v]", got, want)
	}
}

func TestSubgraphClusterOption(t *testing.T) {
	di := NewGraph(Directed)
	sub := di.Subgraph("test", ClusterOption{})
	if got, want := sub.id, "cluster_s0"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEdgeLabel(t *testing.T) {
	di := NewGraph(Directed)
	n1 := di.Node("n1")
	n2 := di.Node("n2")
	n1.Edge(n2, "wat")
	if got, want := flatten(di.String()), `digraph  {node[label="n1"] n1;node[label="n2"] n2;n1->n2[label="wat"];}`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// dot -Tpng cluster.dot > cluster.png && open cluster.png
func TestCluster(t *testing.T) {
	di := NewGraph(Directed)
	outside := di.Node("Outside")
	clusterA := di.Subgraph("Cluster A", ClusterOption{})
	insideOne := clusterA.Node("one")
	insideTwo := clusterA.Node("two")
	clusterB := di.Subgraph("Cluster B", ClusterOption{})
	insideThree := clusterB.Node("three")
	insideFour := clusterB.Node("four")
	outside.Edge(insideFour).Edge(insideOne).Edge(insideTwo).Edge(insideThree).Edge(outside)
	ioutil.WriteFile("doc/cluster.dot", []byte(di.String()), os.ModePerm)
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}
