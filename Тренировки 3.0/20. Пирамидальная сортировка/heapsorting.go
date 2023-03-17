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
		for i := len(h.data) - 1; h.data[(i-1)/2] > h.data[i]; {
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
	for i := 0; len(h.data) > 1 && i*2+1 < len(h.data); {
		if i*2+2 < len(h.data) {
			if h.data[i] >= h.data[i*2+1] || h.data[i] >= h.data[i*2+2] {
				temp := h.data[i]
				if h.data[i*2+1] < h.data[i*2+2] {
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
			if h.data[i] >= h.data[i*2+1] {
				temp := h.data[i]
				h.data[i] = h.data[i*2+1]
				h.data[i*2+1] = temp
				i = i*2 + 1
			} else {
				break
			}
		}
	}

	h.data = h.data[:len(h.data)-1]
	return k
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	count, _ := strconv.Atoi(strings.TrimSpace(string(src)))

	src, _ = reader.ReadBytes('\n')
	src_slice := strings.Split(string(src[:len(src)-1]), " ")

	numbers := &Heap{}

	for _, str := range src_slice {
		num, _ := strconv.Atoi(str)
		numbers.Insert(num)
	}

	fmt.Print(strconv.Itoa(numbers.Extract()))
	for i := 1; i < count; i++ {
		fmt.Print(" " + strconv.Itoa(numbers.Extract()))
	}
	fmt.Print("\n")
}
