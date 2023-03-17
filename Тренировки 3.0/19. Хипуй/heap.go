package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	data []int
}

func (h *Heap) Insert(k int) {
	h.data = append(h.data, k)
	if len(h.data) > 1 {
		for i := len(h.data) - 1; (h.data)[(i-1)/2] < (h.data)[i]; {
			temp := h.data[i]
			h.data[i] = h.data[(i-1)/2]
			h.data[(i-1)/2] = temp
			i = (i - 1) / 2
		}
	}
}

func (h *Heap) Extract() int {
	k := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	if len(h.data) > 1 {
		for i := 0; i <= (len(h.data)-2)/2; {
			if i*2+2 < len(h.data) {
				if h.data[i] < h.data[i*2+1] || h.data[i] < h.data[i*2+2] {
					temp := h.data[i]
					if h.data[i*2+1] > h.data[i*2+2] {
						h.data[i] = h.data[i*2+1]
						h.data[i*2+1] = temp
						i = i*2 + 1
					} else {
						h.data[i] = h.data[i*2+2]
						h.data[i*2+2] = temp
						i = i*2 + 2
					}
				} else {

					break
				}
			} else {
				if h.data[i] < h.data[i*2+1] {
					temp := h.data[i]
					h.data[i] = h.data[i*2+1]
					h.data[i*2+1] = temp
					i = i*2 + 1
				} else {
					break
				}
			}
		}
	}
	h.data = h.data[:len(h.data)-1]
	return k
}

func main() {
	heap := &Heap{}

	reader := bufio.NewReader(os.Stdin)
	commands := []string{}
	src, _, _ := reader.ReadLine()
	num_cmd, _ := strconv.Atoi(string(src))
	for i := 0; i < num_cmd; i++ {
		src, _, _ = reader.ReadLine()
		commands = append(commands, string(src))
	}
	for _, c_string := range commands {
		cmd := strings.Split(c_string, " ")
		if cmd[0] == "0" {
			num, _ := strconv.Atoi(cmd[1])
			heap.Insert(num)
		} else {
			fmt.Println(heap.Extract())
		}
	}
}
