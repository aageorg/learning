package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	data   []int
	cursor int
}

func (q *Queue) Push(i int) {
	q.data = append(q.data, i)
	fmt.Println("ok")

}

func (q *Queue) Front() {
	if len(q.data) == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(strconv.Itoa(q.data[0]))
}

func (q *Queue) Size() {
	fmt.Println(strconv.Itoa(len(q.data)))
}

func (q *Queue) Pull() {
	if len(q.data) == 0 {
		fmt.Println("error")
		return
	}
	var i int
	i = q.data[0]
	q.data = q.data[1:]
	fmt.Println(strconv.Itoa(i))
}

func (q *Queue) Clear() {
	q.data = []int{}
	fmt.Println("ok")
}

func (q *Queue) Do(requests []string) {
	for _, req := range requests {
		command := strings.Split(req, " ")
		switch command[0] {
		case "exit":
			fmt.Println("bye")
			return
		case "push":
			n, _ := strconv.Atoi(command[1])
			q.Push(n)
		case "pop":
			q.Pull()
		case "front":
			q.Front()
		case "clear":
			q.Clear()
		case "size":
			q.Size()
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
	var queue Queue
	queue.Do(reqs)

}
