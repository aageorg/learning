package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Deque struct {
	data []int
	mu   sync.Mutex
}

func (d *Deque) Push(n int, where string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if where == "back" {
		d.data = append(d.data, n)
	} else {
		d.data = append([]int{n}, d.data...)
	}
	fmt.Println("ok")
}

func (d *Deque) Get(fromwhere string, remove bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.data) == 0 {
		fmt.Println("error")
		return
	}
	var position int
	if fromwhere == "back" {
		position = len(d.data) - 1
	} else {
		position = 0
	}
	n := d.data[position]
	if remove {
		d.data = append(d.data[:position], d.data[position+1:]...)
	}
	fmt.Println(n)
}

func (d *Deque) Size() {
	d.mu.Lock()
	defer d.mu.Unlock()
	fmt.Println(len(d.data))
}

func (d *Deque) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data = d.data[:0]
	fmt.Println("ok")

}

func (d *Deque) Do(requests []string) {
	for _, req := range requests {
		command := strings.Split(req, " ")
		switch command[0] {
		case "exit":
			fmt.Println("bye")
			return
		case "push_front":
			n, _ := strconv.Atoi(command[1])
			d.Push(n, "front")
		case "push_back":
			n, _ := strconv.Atoi(command[1])
			d.Push(n, "back")
		case "pop_front":
			d.Get("front", true)
		case "pop_back":
			d.Get("back", true)
		case "front":
			d.Get("front", false)
		case "back":
			d.Get("back", false)
		case "size":
			d.Size()
		case "clear":
			d.Clear()

		}

	}
}

func main() {
	var reqs []string
	reader := bufio.NewReader(os.Stdin)
	for src, _, _ := reader.ReadLine(); string(src) != "exit"; {
		reqs = append(reqs, string(src))
		src, _, _ = reader.ReadLine()
	}
	reqs = append(reqs, "exit")
	var deque Deque
	deque.Do(reqs)

}
