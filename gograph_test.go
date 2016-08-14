package gograph

import (
    "testing"
    "gograph/class"
    "gograph/alg"
    "fmt"
    "strings"
)

func TestGographDirected(t *testing.T){
    g := &graph.DiGraph{}
    g.AddNode(1).AddNode(2).AddNode(3)
    g.AddEdge(1, 2, 0).AddEdge(1, 3, 1)
    // String() must recived a pointer
    fmt.Printf("%s\n", g)
    fmt.Println(strings.Join(g.Edges(), "\n"))
    fmt.Println(alg.HasSimplePath(g, 1, 2))
    fmt.Println(alg.HasSimplePath(g, 2, 2))
    fmt.Println(alg.HasSimplePath(g, 3, 4))
}


func TestAlgShortestPath( t *testing.T){
    g := &graph.DiGraph{}
    g.AddNode(1).AddNode(2).AddNode(3)
    g.AddEdge(1, 2, 2).AddEdge(1, 3, 1).AddEdge(2, 3, 1)
    path, d , err := alg.ShortestPath(g, 1, 3)
    
    if err!= nil {t.Fatal()}
    if d!=1{t.Error("The ShortestPath is wrong")}
    fmt.Printf("%d\n", len(path))

    g.AddEdge(2, 4, 4).AddEdge(3, 4, 1)
    path, d , err = alg.ShortestPath(g, 1, 4)
    if d!=2 {t.Error("The ShortestPath is wrong")}
}