package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Calc(a int, b int, op byte) int {
	switch op {
	case '*':
		return a * b
	case '+':
		if b == 1 {
			a++
			return a
		}
		return a + b
	case '-':
		if b == 1 {
			a--
			return a
		}
		return a - b
	}
	return -1
}

type Stack struct {
	data [][2]int
}

func (s *Stack) Push(b int) {
	if len(s.data) > 0 && s.data[0][0] == b {
		s.data[0][1]++
	} else {
		s.data = append([][2]int{{b, 1}}, s.data...)
	}
}

func (s *Stack) Pull() (b int) {
	if len(s.data) > 0 {
		b = s.data[0][0]
		s.data[0][1]--
		if s.data[0][1] == 0 {
			s.data = s.data[1:]
		}
	}
	return
}

func main() {
	s := Stack{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	src, _ := reader.ReadBytes(' ')
	a, _ := strconv.Atoi(string(src[:len(src)-1]))
	s.Push(a)
	src, _ = reader.ReadBytes(' ')
	b, _ := strconv.Atoi(string(src[:len(src)-1]))
	for {
		src, _ = reader.ReadBytes(' ')
		if len(src) == 1 {
			continue
		}
		next, err := strconv.Atoi(string(src[:len(src)-1]))
		if err == nil {
			s.Push(b)
			b = next
		} else {
			a = s.Pull()
			b = Calc(a, b, src[0])
		}
		if src[len(src)-1] == '\n' {
			break
		}
	}
	fmt.Println(b)

}
