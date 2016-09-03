package alg

import (
    "testing"
    "gograph/graph"
    "fmt"
)

func TestTopoOder( t *testing.T){
    g:= &graph.DiGraph{}
    g.AddNode(1).AddNode(2).AddNode(3)
    g.AddEdge(1, 2, 0).AddEdge(1, 3, 1)
    fmt.Printf("The topo order is: %v\n", TopoOrder(g))
}