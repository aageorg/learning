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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	nails := &Heap{}

	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	count, _ := strconv.Atoi(strings.TrimSpace(string(src)))
	src, _, _ = reader.ReadLine()
	nails_src := strings.Split(strings.TrimSpace(string(src)), " ")

	for _, nail := range nails_src {
		n, _ := strconv.Atoi(nail)
		nails.Insert(n)
	}

	a := nails.Extract()
	b := nails.Extract()

	dp := make([]int, count)
	for i := 0; i < count-1; i++ {
		if i == 0 {
			dp[i] = b - a
			continue
		}
		if i == 1 {
			dp[i] = b - a
			a = b
			b = nails.Extract()
			continue
		}
		if i == 2 {
			dp[i] = dp[i-1] + (b - a)
			a = b
			b = nails.Extract()
			continue
		}

		dp[i] = Min(dp[i-1], dp[i-2]) + (b - a)
		a = b
		b = nails.Extract()
	}
	if count == 2 {
		fmt.Println(b - a)
		return
	}
	dp[count-1] = Min(dp[count-2], dp[count-3]) + (b - a)
	fmt.Println(dp[count-1])
}
