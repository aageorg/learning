package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Fibo(a int, k int, dp []int) int {
	if dp[a] == 0 {
		for i := a; i >= a-k && i > -1; i-- {
			dp[a] += dp[i]
		}
	}
	return dp[a]
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	input_pair := strings.Split(strings.TrimSpace(string(src)), " ")

	N, _ := strconv.Atoi(input_pair[0])
	k, _ := strconv.Atoi(input_pair[1])

	if N <= 2 {
		fmt.Println(1)
		return
	}
	dp := make([]int, N)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < N; i++ {
		Fibo(i, k, dp)
	}
	fmt.Println(dp[N-1])
}
