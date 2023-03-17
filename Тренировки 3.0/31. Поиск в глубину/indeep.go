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

type Heap struct {
	data []int
}

func (h *Heap) Insert(k int) {
	h.data = append(h.data, k)
	if len(h.data) > 1 {
		for i := len(h.data) - 1; h.data[(i-1)/2] > h.data[i]; {
			temp := h.data[i]
			h.data[i] = h.data[(i-1)/2]
			h.data[(i-1)/2] = temp
			i = (i - 1) / 2
		}
	}
}

func (h *Heap) Extract() int {
	k := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	for i := 0; len(h.data) > 1 && i*2+1 < len(h.data); {
		if i*2+2 < len(h.data) {
			if h.data[i] >= h.data[i*2+1] || h.data[i] >= h.data[i*2+2] {
				temp := h.data[i]
				if h.data[i*2+1] < h.data[i*2+2] {
					h.data[i] = h.data[i*2+1]
					h.data[i*2+1] = temp
					i = i*2 + 1
				} else {
					h.data[i] = h.data[i*2+2]
					h.data[i*2+2] = temp
					i = i*2 + 2
				}
			} else {
				break
			}
		} else {
			if h.data[i] >= h.data[i*2+1] {
				temp := h.data[i]
				h.data[i] = h.data[i*2+1]
				h.data[i*2+1] = temp
				i = i*2 + 1
			} else {
				break
			}
		}
	}

	h.data = h.data[:len(h.data)-1]
	return k
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
		if from != to {
			gr.ver[to].neigh = append(gr.ver[to].neigh, gr.ver[from])
		}
	}
	r := gr.SearchComponents()
	fmt.Println(len(r[0].ver))
	var h Heap
	for i := 0; i < len(r[0].ver); i++ {
		h.Insert(r[0].ver[i].id)
	}
	fmt.Print(strconv.Itoa(h.Extract()))
	for i := 1; i < len(r[0].ver); i++ {
		fmt.Print(" " + strconv.Itoa(h.Extract()))
	}
	fmt.Print("\n")

}
