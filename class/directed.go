package graph
import (
    "fmt"
    "errors"
)

type Vertex interface{}

// the int32 with each edge is the weight of the edge 
type aDjency map[Vertex]int32

type DiGraph struct{
    vertices map[Vertex]aDjency
}

type DiGraphInterface interface{
    AddEdge(v1, v2 Vertex, weight int32) *DiGraph
    AddNode(v Vertex) *DiGraph
    String() string
    Edges() []string
    Len() int
    Nodes() []Vertex
    Adjs(Vertex) []Vertex
    HasEdge(v1, v2 Vertex) bool
    HasNode(Vertex) bool
    Weight(v1, v2 Vertex) (int32, error)
}

func (d *DiGraph)AddEdge(v1, v2 Vertex, weight int32)(*DiGraph){
    d.AddNode(v1).AddNode(v2)
    d.vertices[v1][v2] = weight
    return d
}

func (d *DiGraph) AddNode(v Vertex)(*DiGraph){
    if d.vertices == nil {d.vertices =make(map[Vertex]aDjency)}
    _, exist := d.vertices[v]
    if !exist {
        d.vertices[v] = aDjency{}
    }
    return d
}

func (d *DiGraph)String() string {
    var nodes []Vertex
    for k, _ := range d.vertices{
        nodes = append(nodes, k)
    }
    return fmt.Sprintf("DiGraph with nodes: %v", nodes)
}

func (d *DiGraph)Edges() []string{
    var  edges []string
    for k, v := range d.vertices{
        for adj, weight := range v{
            edges = append(edges, fmt.Sprintf("%d %d %d", k, adj, weight ))
        }
    }
    return edges
}


func (d *DiGraph)Len() int{
    return len(d.vertices)
}

func (d *DiGraph)Nodes() (nodes []Vertex){
    for k, _ := range d.vertices{
        nodes = append(nodes, k)
    }
    return 
}

func (d *DiGraph)Adjs(v Vertex) (nodes []Vertex){
    adj := d.vertices[v]
    for node, _ := range adj {
        nodes = append(nodes, node)
    }
    return
}

func (d *DiGraph)HasNode(v Vertex) bool {
    _, exist := d.vertices[v]
    return exist
}

func (d *DiGraph)HasEdge(v1, v2 Vertex) bool{
    if d.HasNode(v1) && d.HasNode(v2){
        _, exist := d.vertices[v1][v2]
        return exist
    }
    return false
}

func (d *DiGraph)Weight(v1, v2 Vertex)(int32, error){
    if d.HasEdge(v1, v2){
        return d.vertices[v1][v2], nil
    }
    return 0, errors.New("Does not has this edge")
}