package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Queue struct {
	data []int
}

func (qu *Queue) Len() int {
	return len(qu.data)
}

func (qu *Queue) Add(elem int) {
	qu.data = append(qu.data, elem)
}

func (qu *Queue) Get() (first int) {
	first = qu.data[0]
	qu.data = qu.data[1:]
	return
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
	mu    sync.Mutex
}

func (v *Vertex) SetLabel(n int) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.label = n
}

func (v *Vertex) Label() int {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.label
}

func (gr *Graph) Reset() {
	for i := 1; i < len(gr.ver); i++ {
		gr.ver[i].SetLabel(0)
	}
}

func (gr *Graph) SearchMinWay(a, b int) (int, []int) {
	way := make(map[int]int)
	way[a] = 0
	gr.ver[a].SetLabel(0)
	qu := &Queue{}
	qu.Add(a)
	for qu.Len() > 0 {
		id := qu.Get()
		if gr.ver[id].id == b {
			result := []int{b}
			for i := b; i != a; {
				result = append([]int{way[i]}, result...)
				i = way[i]
			}
			return gr.ver[id].Label(), result
		}
		for i, _ := range gr.ver[id].neigh {
			if gr.ver[id].neigh[i].Label() == -1 {
				way[gr.ver[id].neigh[i].id] = gr.ver[id].id
				gr.ver[id].neigh[i].SetLabel(gr.ver[id].Label() + 1)
				qu.Add(gr.ver[id].neigh[i].id)
			}
		}
	}
	return -1, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(string(src))

	gr := Graph{ver: make([]*Vertex, num+1)}
	for i := 1; i < len(gr.ver); i++ {
		v := &Vertex{id: i, label: -1}
		gr.ver[i] = v
	}
	for i := 1; i < num+1; i++ {
		src, _, _ = reader.ReadLine()
		vsx := strings.Split(strings.TrimSpace(string(src)), " ")
		for j, e := range vsx {
			if e == "1" {
				gr.ver[i].neigh = append(gr.ver[i].neigh, gr.ver[j+1])
			}
		}

	}
	src, _, _ = reader.ReadLine()
	se_src := strings.Split(strings.TrimSpace(string(src)), " ")
	start, _ := strconv.Atoi(se_src[0])
	end, _ := strconv.Atoi(se_src[1])
	steps, way := gr.SearchMinWay(start, end)
	fmt.Println(steps)
	if len(way) > 1 {
		fmt.Print(strconv.Itoa(way[0]))
		for i := 1; i < len(way); i++ {
			fmt.Print(" " + strconv.Itoa(way[i]))
		}
		fmt.Print("\n")
	}
}
