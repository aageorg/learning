package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(string(src))

	dp := map[int]int{0: 1, 1: 2, 2: 4, 3: 7}
	for i := 4; i <= num; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	fmt.Println(dp[num])

}
