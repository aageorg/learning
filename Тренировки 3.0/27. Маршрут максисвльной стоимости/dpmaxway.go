package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func Max(a, b int) (int, int) {
	if a > b {
		return a, 0
	}
	return b, 1
}

type Table struct {
	field [][]int
}

func makeTable(x int, y int) Table {
	t := Table{field: make([][]int, x+1)}
	for i := 0; i < len(t.field); i++ {
		t.field[i] = make([]int, y+1)
	}
	return t
}

// 0 - Right; 1 - Bottom

func main() {
	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	xy := bytes.Split(src, []byte{' '})
	x, _ := strconv.Atoi(string(xy[0]))
	y, _ := strconv.Atoi(string(xy[1]))
	table := makeTable(x, y)

	for i := 1; i <= x; i++ {
		src, _, _ = reader.ReadLine()
		xy := bytes.Split(src, []byte{' '})
		for ind, num := range xy {
			table.field[i][ind+1], _ = strconv.Atoi(string(num))
		}

	}
	dp := make(map[[2]int][3]int, x*y)
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			max, dir := Max(table.field[i-1][j], table.field[i][j-1])
			if i == 1 {
				dir = 1
			}
			if j == 1 {
				dir = 0
			}
			table.field[i][j] = table.field[i][j] + max
			var a, b int
			if dir == 0 {
				a = i - 1
				b = j
			} else {
				a = i
				b = j - 1
			}
			dp[[2]int{i, j}] = [3]int{a, b, dir}
		}
	}

	dirtostr := func(i int) string {
		if i == 0 {
			return "D"
		} else {
			return "R"
		}
	}
	fmt.Println(table.field[x][y])
	output := dirtostr(dp[[2]int{x, y}][2])
	if len(dp) == 1 {
		return
	}
	for i, j := dp[([2]int{x, y})][0], dp[([2]int{x, y})][1]; i > 1 || j > 1; {
		output = dirtostr(dp[([2]int{i, j})][2]) + " " + output
		i, j = dp[([2]int{i, j})][0], dp[([2]int{i, j})][1]
	}
	fmt.Println(output)

}
