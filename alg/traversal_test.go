package alg

import (
    "testing"
    "gograph/graph"
    "fmt"
)

func TestTraversel(t *testing.T){
    g:= &graph.DiGraph{}
    g.AddNode(1).AddNode(2).AddNode(3)
    g.AddEdge(1, 2, 0).AddEdge(1, 3, 1)
    g.AddEdge(1, 4, 9)
    print := func(node graph.Vertex){
        fmt.Printf("node %v\n", node)
    }
    fmt.Println("DFS")
    DFS(g, 1, print)
    fmt.Println("BFS")
    BFS(g, 1, print)
}