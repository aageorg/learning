package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(i int) {
	s.data = append([]int{i}, s.data...)
}

func (s *Stack) Pull() int {
	var i int
	if len(s.data) > 0 {
		i = s.data[0]
		s.data = s.data[1:]
	}
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

func dfs(now *Vertex, prev *Vertex, dp *Stack) bool {
	now.SetLabel(1)
	dp.Push(now.id)

	for i, _ := range now.neigh {
		if now.neigh[i] == prev || now.neigh[i].Label() == 2 {
			continue
		}
		if now.neigh[i].Label() == 0 {
			if dfs(now.neigh[i], now, dp) {
				return true
			}
		} else {
			dp.Push(now.neigh[i].id)
			return true
		}
	}
	now.SetLabel(2)
	_ = dp.Pull()
	return false
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

func (gr *Graph) SearchCycle() (bool, []int) {
	for i := 1; i < len(gr.ver); i++ {
		gr.Reset()
		dp := Stack{}
		if dfs(gr.ver[i], nil, &dp) {
			var result []int
			entry := dp.Pull()
			for j := dp.Pull(); j != entry; {
				result = append(result, j)
				j = dp.Pull()
			}
			result = append(result, entry)
			return true, result
		}

	}
	return false, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(string(src))

	gr := Graph{ver: make([]*Vertex, num+1)}
	for i := 1; i < len(gr.ver); i++ {
		v := &Vertex{id: i}
		gr.ver[i] = v
	}
	for i := 1; i < num+1; i++ {
		for j := 1; j < num+1; j++ {
			matrix := make([]byte, 2)
			reader.Read(matrix)
			if matrix[0] == '1' {
				gr.ver[i].neigh = append(gr.ver[i].neigh, gr.ver[j])
			}
			if matrix[1] == '\n' {
				break
			}
		}
	}
	found, cycle := gr.SearchCycle()
	if found {
		fmt.Println("YES")
		fmt.Println(len(cycle))
		fmt.Print(strconv.Itoa(cycle[0]))
		for i := 1; i < len(cycle); i++ {
			fmt.Print(" " + strconv.Itoa(cycle[i]))
		}
		fmt.Print("\n")
	} else {
		fmt.Println("NO")
	}
}
