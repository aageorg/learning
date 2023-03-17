package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var N int
var M int

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
	strs int
	cols int
}

type Vertex struct {
	id    int
	neigh []*Vertex
	label int
	str   int
	col   int
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

func (v *Vertex) GetNeighs(gr *Graph) {
	neigh_map := [8][2]int{[2]int{2, 1}, [2]int{1, 2}, [2]int{-1, 2}, [2]int{-2, 1}, [2]int{-2, -1}, [2]int{-1, -2}, [2]int{1, -2}, [2]int{2, -1}}

	for _, xy := range neigh_map {

		if v.str+xy[0] > 0 && v.str+xy[0] <= N && v.col+xy[1] > 0 && v.col+xy[1] <= M && !v.HasNeigh(M*(v.str+xy[0]-1)+v.col+xy[1]) {
			if gr.ver[M*(v.str+xy[0]-1)+v.col+xy[1]] == nil {
				var Vtx = Vertex{id: M*(v.str+xy[0]-1) + v.col + xy[1], str: v.str + xy[0], col: v.col + xy[1], label: -1}
				v.neigh = append(v.neigh, &Vtx)
			} else {
				v.neigh = append(v.neigh, gr.ver[M*(v.str+xy[0]-1)+v.col+xy[1]])

			}
		}
	}
}

func (v *Vertex) HasNeigh(id int) bool {
	for i, _ := range v.neigh {
		if v.neigh[i].id == id {
			return true
		}
	}
	return false
}

func (gr *Graph) Reset() {
	for i := 1; i < len(gr.ver); i++ {
		if gr.ver[i] != nil {
			gr.ver[i].SetLabel(-1)
		}
	}
}

func (gr *Graph) SearchMinWay(a int, fleas map[int]int) int {

	var sum int
	gr.ver[a].SetLabel(0)
	qu := &Queue{}
	qu.Add(a)
	for qu.Len() > 0 {
		id := qu.Get()
		if _, ok := fleas[gr.ver[id].id]; ok {
			sum += gr.ver[id].Label()
			delete(fleas, gr.ver[id].id)
		}
		gr.ver[id].GetNeighs(gr)
		for i, _ := range gr.ver[id].neigh {
			if gr.ver[gr.ver[id].neigh[i].id] == nil {
				gr.ver[gr.ver[id].neigh[i].id] = gr.ver[id].neigh[i]
			}
			if gr.ver[id].neigh[i].Label() == -1 {
				gr.ver[id].neigh[i].SetLabel(gr.ver[id].Label() + 1)
				qu.Add(gr.ver[id].neigh[i].id)
				if _, ok := fleas[gr.ver[id].neigh[i].id]; ok {
					sum += gr.ver[id].neigh[i].Label()
					delete(fleas, gr.ver[id].neigh[i].id)
				}

			}
		}
	}
	if len(fleas) == 0 {
		return sum
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	src, _, _ := reader.ReadLine()
	se_src := strings.Split(strings.TrimSpace(string(src)), " ")
	N, _ = strconv.Atoi(se_src[0])
	M, _ = strconv.Atoi(se_src[1])
	S, _ := strconv.Atoi(se_src[2])
	T, _ := strconv.Atoi(se_src[3])
	Q, _ := strconv.Atoi(se_src[4])
	gr := Graph{ver: make([]*Vertex, N*M+1), strs: N, cols: M}
	gr.ver[M*(S-1)+T] = &Vertex{id: M*(S-1) + T, str: S, col: T, label: -1}
	var fleas = make(map[int]int)
	for i := 0; i < Q; i++ {
		src, _, _ := reader.ReadLine()
		flea := strings.Split(strings.TrimSpace(string(src)), " ")
		y, _ := strconv.Atoi(flea[0])
		x, _ := strconv.Atoi(flea[1])
		fleas[M*(y-1)+x] = -1
		gr.ver[M*(y-1)+x] = &Vertex{id: M*(y-1) + x, str: y, col: x, label: -1}
	}
	sum := gr.SearchMinWay(M*(S-1)+T, fleas)
	fmt.Println(sum)
}
