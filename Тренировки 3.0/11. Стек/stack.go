package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Stack struct {
	mu     sync.Mutex
	data   []int
	cursor int
}

func makeStack(length int) Stack {
	return Stack{data: make([]int, length), cursor: length - 1}
}

func (s *Stack) Push(i int) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[s.cursor] = i
	s.cursor--
	if len(s.data)/s.cursor > 4 {
		temp := len(s.data)
		s.data = append(make([]int, temp/4), s.data...)
		s.cursor = s.cursor + temp/4
	}
	return "ok"
}

func (s *Stack) Back() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cursor == len(s.data)-1 {
		return "error"
	}
	return strconv.Itoa(s.data[s.cursor+1])
}

func (s *Stack) Size() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return strconv.Itoa(len(s.data) - s.cursor - 1)
}

func (s *Stack) Pull() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cursor == len(s.data)-1 {
		return "error"
	}
	var i int
	s.cursor++
	i = s.data[s.cursor]
	s.data[s.cursor] = 0
	if len(s.data)/s.cursor < 2 && len(s.data) > 40 {
		temp := len(s.data)
		s.data = s.data[len(s.data)/4:]
		s.cursor = s.cursor - temp/4

	}
	return strconv.Itoa(i)
}

func (s *Stack) Clear() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make([]int, 40)
	s.cursor = 39
	return "ok"
}

func main() {
	var wg sync.WaitGroup
	var stack = makeStack(40)
	reqs := make(chan string, 20)
	var output string
	wg.Add(1)
	go func() {
		defer wg.Done()
		for req := range reqs {
			if req == "exit" {
				output += "bye\n"
				return
			}
			command := strings.Split(req, " ")
			switch command[0] {
			case "push":
				n, _ := strconv.Atoi(command[1])
				output += stack.Push(n) + "\n"
			case "pop":
				output += stack.Pull() + "\n"
			case "back":
				output += stack.Back() + "\n"
			case "clear":
				output += stack.Clear() + "\n"
			case "size":
				output += stack.Size() + "\n"
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for src, _, _ := reader.ReadLine(); string(src) != "exit"; {
		reqs <- string(src)
		src, _, _ = reader.ReadLine()
	}
	reqs <- "exit"
	wg.Wait()
	fmt.Print(output)

}
