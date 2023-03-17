package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	data   []int
	cursor int
}

func makeStack(length int) Stack {
	return Stack{data: make([]int, length), cursor: length - 1}
}

func (s *Stack) Push(i int) {
	if s.cursor == -1 {
		return
	}
	s.data[s.cursor] = i
	s.cursor--
}

func (s *Stack) Pull() int {
	var i int
	if s.cursor == len(s.data)-1 {
		return 0
	}
	s.cursor++
	i = s.data[s.cursor]
	s.data[s.cursor] = 0
	return i
}

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

func dfs(now *Vertex, dp *Stack) bool {
	now.SetLabel(1)
	for i, _ := range now.neigh {
		if now.neigh[i].Label() == 2 {
			continue
		}
		if now.neigh[i].Label() == 0 {
			if !dfs(now.neigh[i], dp) {
				return false
			}
		} else {
			return false
		}
	}
	now.SetLabel(2)
	dp.Push(now.id)
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

func (gr *Graph) Topology() (bool, []int) {
	dp := makeStack(len(gr.ver) - 1)
	for i := 1; i < len(gr.ver); i++ {
		if gr.ver[i].Label() == 2 {
			continue
		}
		if !dfs(gr.ver[i], &dp) {
			return false, nil
		}
	}
	return true, dp.data
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
	}
	found, topo := gr.Topology()
	if found {
		fmt.Print(strconv.Itoa(topo[0]))
		for i := 1; i < len(topo); i++ {
			fmt.Print(" " + strconv.Itoa(topo[i]))
		}

		fmt.Print("\n")
	} else {
		fmt.Println("-1")
	}
}
