package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	ver  []*Vertex
	now  *Vertex
	prev *Vertex
}

type Vertex struct {
	id    int
	neigh []*Vertex
	label int
}

func dfs(now *Vertex, l int) bool {
	now.SetLabel(l)
	for i, _ := range now.neigh {
		if now.neigh[i].Label() == 3-l {
			continue
		}
		if now.neigh[i].Label() == l {
			return false
		}
		if now.neigh[i].Label() == 0 {
			if !dfs(now.neigh[i], 3-l) {
				return false
			}
		}
	}
	return true
}

func (v *Vertex) SetLabel(n int) {
	v.label = n
}

func (v *Vertex) Label() int {
	return v.label
}

func (gr *Graph) Reset() {
	for i := 1; i < len(gr.ver); i++ {
		gr.ver[i].SetLabel(0)
	}
}

func (gr *Graph) IsDual() bool {
	for i := 1; i < len(gr.ver); i++ {
		if gr.ver[i].Label() == 0 {
			if !dfs(gr.ver[i], 1) {
				return false
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	veredges := bytes.Split(src, []byte(" "))
	vert_num, _ := strconv.Atoi(string(veredges[0]))
	edge_num, _ := strconv.Atoi(string(veredges[1]))
	gr := Graph{ver: make([]*Vertex, vert_num+1)}
	for i := 1; i < len(gr.ver); i++ {
		v := &Vertex{id: i}
		gr.ver[i] = v
	}
	for i := 0; i < edge_num; i++ {
		src, _, _ := reader.ReadLine()
		edge := bytes.Split(src, []byte(" "))
		from, _ := strconv.Atoi(string(edge[0]))
		to, _ := strconv.Atoi(string(edge[1]))
		gr.ver[from].neigh = append(gr.ver[from].neigh, gr.ver[to])
		gr.ver[to].neigh = append(gr.ver[to].neigh, gr.ver[from])
	}
	if gr.IsDual() {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
