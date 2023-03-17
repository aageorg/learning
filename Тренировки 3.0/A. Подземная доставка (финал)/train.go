package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var workers map[string]*Worker

type Worker struct {
	counter int
}

func (w *Worker) Add(num int) {
	w.counter = w.counter + num
}

func (w *Worker) Delete(num int) {
	w.counter = w.counter - num
}

func (w *Worker) Sum() int {
	return w.counter
}

type Stack struct {
	mu     sync.Mutex
	names  []string
	ints   []int
	cursor int
}

func makeStack(length int) Stack {
	return Stack{names: make([]string, length), ints: make([]int, length), cursor: length - 1}
}

func (s *Stack) Push(st string, i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.names[s.cursor] = st
	s.ints[s.cursor] = i
	s.cursor--
	if len(s.names)/s.cursor > 4 {
		temp := len(s.names)
		s.names = append(make([]string, temp/4), s.names...)
		s.ints = append(make([]int, temp/4), s.ints...)
		s.cursor = s.cursor + temp/4
	}
}

func (s *Stack) Back() (string, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cursor == len(s.names)-1 {
		return "error", 0
	}
	return s.names[s.cursor+1], s.ints[s.cursor+1]
}

func (s *Stack) Pull() (string, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cursor == len(s.names)-1 {
		return "error", 0
	}
	var st string
	var i int
	s.cursor++
	i = s.ints[s.cursor]
	s.ints[s.cursor] = 0
	st = s.names[s.cursor]
	s.names[s.cursor] = ""
	if len(s.names)/s.cursor < 2 && len(s.names) > 40 {
		temp := len(s.names)
		s.names = s.names[len(s.names)/4:]
		s.ints = s.ints[len(s.ints)/4:]
		s.cursor = s.cursor - temp/4

	}
	return st, i
}

func (s *Stack) Do(requests []string) {
	for _, req := range requests {
		command := strings.Split(req, " ")
		switch command[0] {
		case "add":
			n, _ := strconv.Atoi(command[1])
			s.Push(command[2], n)
			if _, ok := workers[command[2]]; !ok {
				workers[command[2]] = &Worker{}
			}
			workers[command[2]].Add(n)
		case "get":
			if _, ok := workers[command[1]]; !ok {
				fmt.Println(0)
			} else {
				fmt.Println(strconv.Itoa(workers[command[1]].Sum()))
			}
		case "delete":
			n, _ := strconv.Atoi(command[1])
			name, count := s.Pull()
			for count < n {
				workers[name].Delete(count)
				n = n - count
				name, count = s.Pull()
			}
			workers[name].Delete(n)
			if count-n > 0 {
				s.Push(name, count-n)
			}
		}
	}
}

func main() {
	stack := makeStack(40)
	workers = make(map[string]*Worker)
	var reqs []string
	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(strings.TrimSpace(string(src)))
	for i := 0; i < num; i++ {
		src, _, _ = reader.ReadLine()
		reqs = append(reqs, strings.TrimSpace(string(src)))
	}
	stack.Do(reqs)
}
