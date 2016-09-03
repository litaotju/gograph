package alg

import (
    "gograph/graph"
    "container/list"
    "errors"
    "fmt"
)

//深度优先遍历的方法判断图中两点是否有路径相连
func HasSimplePath(this *graph.DiGraph,  v1, v2 graph.Vertex) bool{
    if ! (this.HasNode(v1) && this.HasNode(v2)){
        // fmt.Println("One node is not even in the graph")
        return false
    }
    visited := make(map[graph.Vertex]bool)
    for _, v := range this.Nodes(){
        visited[v] = false
    }
    return dfs(this, v1, v2, visited)
}

func dfs(this *graph.DiGraph, v, target graph.Vertex,  visited map[graph.Vertex]bool) bool{
    find := false
    for _, next := range this.Adjs(v) {
        if vis := visited[next]; !vis{
            visited[next] = true
            if next == target{
                return true
            }
            find = dfs(this, next, target, visited)
            if find{
                return true
            }
        }
    }
    return find
}

// Bell-Ford 方法求最短路径
func ShortestPath(this *graph.DiGraph, s, t graph.Vertex)(nodes []graph.Vertex, dis int32, err error){
    distance := make(map[graph.Vertex]int32)
    parent := make(map[graph.Vertex]graph.Vertex)
    var  MAX int32 = 10000000
    for _, v := range this.Nodes(){
        distance[v] = MAX
        parent[v] = nil
    }
    distance[s] = 0
    for i:=0; i < this.Len()-1; i++{
        for _, u := range this.Nodes(){
            for _, v := range this.Adjs(u){
                 duv, _ := this.Weight(u, v)
                 duv += distance[u]
                 dv :=  distance[v]
                 if duv < dv{ distance[v] = duv; parent[v] = u;}
            }
        }
    }

    result := list.New()
    pre := parent[t]
    for ; pre != nil; pre = parent[pre]{
        result.PushFront(pre)
        if pre == s{
            break
        }
    }
    if pre == nil{
        return nodes, 0, errors.New(fmt.Sprintf("Do not has a connect from %v to %v", s, t))
    }
    for e:= result.Front(); e!=nil ; e = e.Next(){
        nodes = append(nodes, e)
    }
    return nodes, distance[t], nil
}
