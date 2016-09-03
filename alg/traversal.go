package alg

import (
    "gograph/graph"
    "container/list"
    "fmt"
)


func DFS(this *graph.DiGraph, source graph.Vertex, action func(graph.Vertex) ){
    visited := make(map[graph.Vertex]bool)
    for _, v := range this.Nodes(){
        visited[v] = false
    }
    visited[source] = true
    action(source)
    dfsImpl(this, source, visited, action)
}

func dfsImpl(this *graph.DiGraph, v graph.Vertex, visited map[graph.Vertex]bool, action func(graph.Vertex) ){
    for _, next := range this.Adjs(v) {
        if vis := visited[next]; !vis{
            visited[next] = true
            action(next)
            dfsImpl(this, next, visited, action)
        }
    }
}

// TODO: 完成广度优先搜索
func BFS(this *graph.DiGraph, source graph.Vertex, action func(graph.Vertex) ){
    visited := make(map[graph.Vertex]bool)
    for _, v := range this.Nodes(){
        visited[v] = false
    }
    visited[source] = true
    action(source)
    queue := list.New()
    queue.PushBack(source)
    for ;queue.Len() > 0; {
        head := graph.Vertex(queue.Front())
        for _, adj := range this.Adjs(head){
            if vis := visited[adj]; !vis {
                visited[adj] = true
                queue.PushBack(adj)
                action(adj)
            }
        }
        queue.Remove(queue.Front())
    }
}



