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

func dfs(now *Vertex, g *Graph) {
	now.SetLabel(1)
	g.ver = append(g.ver, now)
	for i := 0; i < len(now.neigh); i++ {
		if now.neigh[i].Label() == 0 {
			dfs(now.neigh[i], g)
		}
	}
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

func (gr *Graph) SearchComponents() []Graph {
	var result []Graph
	for i := 1; i < len(gr.ver); i++ {
		if gr.ver[i].Label() == 0 {
			g := Graph{}
			dfs(gr.ver[i], &g)
			result = append(result, g)
		}

	}
	return result
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
	r := gr.SearchComponents()
	fmt.Println(len(r))
	for _, comp := range r {
		fmt.Println(len(comp.ver))
		fmt.Print(strconv.Itoa(comp.ver[0].id))
		for i := 1; i < len(comp.ver); i++ {
			fmt.Print(" " + strconv.Itoa(comp.ver[i].id))
		}
		fmt.Print("\n")
	}
}
