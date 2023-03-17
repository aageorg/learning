package main

import (
	"bufio"
	"fmt"
	"os"
)

var Schema = map[byte]byte{'(': ')', '{': '}', '[': ']'}

type Stack struct {
	data []byte
}

func (s *Stack) Push(b byte) {
	s.data = append([]byte{b}, s.data...)
}

func (s *Stack) Pull() byte {
	var b byte
	if len(s.data) > 0 {
		b = s.data[0]
		s.data = s.data[1:]
	}
	return b
}

func main() {
	s := Stack{}

	reader := bufio.NewReader(os.Stdin)
	src, _ := reader.ReadBytes('\n')

	input := src[:len(src)-1]

	for _, b := range input {
		if _, ok := Schema[b]; ok {
			s.Push(b)
		} else {
			open := s.Pull()
			if Schema[open] != b {
				fmt.Println("no")
				return
			}
		}
	}
	if len(s.data) == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
