package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DP struct {
	num int
	m   map[int][2]int
}

func (dp *DP) Restore() (result []int) {
	result = []int{1}
	for i := 1; i < dp.num; {
		result = append(result, dp.m[i][1])
		i = dp.m[i][1]
	}
	return
}

func F(n int, s int, dp *DP) {
	if n == 1 {
		return
	}

	if n%3 == 0 {
		if res, ok := dp.m[n/3]; ok && s+1 >= res[0] {
			return
		} else {
			dp.m[n/3] = [2]int{s + 1, n}
			F(n/3, s+1, dp)
		}
	}

	if n%2 == 0 {
		if res, ok := dp.m[n/2]; ok && s+1 >= res[0] {
			return
		} else {
			dp.m[n/2] = [2]int{s + 1, n}
			F(n/2, s+1, dp)
		}
	}

	if res, ok := dp.m[n-1]; ok && s+1 >= res[0] {
		return
	} else {
		dp.m[n-1] = [2]int{s + 1, n}
		F(n-1, s+1, dp)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(strings.TrimSpace(string(src)))
	dp := DP{num: n, m: map[int][2]int{n: [2]int{0, 0}}}
	F(n, 0, &dp)
	r := dp.Restore()
	fmt.Println(len(r) - 1)
	fmt.Print(strconv.Itoa(r[0]))
	for i := 1; i < len(r); i++ {
		fmt.Print(" " + strconv.Itoa(r[i]))
	}
	fmt.Print("\n")
}
