package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Min(a, b float64) (float64, int) {
	if a < b {
		return a, 0
	}
	return b, 1
}

type Table struct {
	field [][]float64
}

func makeTable(x int, y int) Table {
	t := Table{field: make([][]float64, x+1)}
	for i := 0; i < len(t.field); i++ {
		t.field[i] = make([]float64, y+1)
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

	for i := 0; i <= x; i++ {
		if i == 0 {
			for j := 0; j < len(table.field[i]); j++ {
				table.field[i][j] = math.Inf(+1)
			}
			continue
		}
		if i > 1 {
			table.field[i][0] = math.Inf(+1)
		}

		src, _, _ = reader.ReadLine()
		xy := bytes.Split(src, []byte{' '})
		for ind, num := range xy {
			table.field[i][ind+1], _ = strconv.ParseFloat(string(num), 64)
		}

	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			min, _ := Min(table.field[i-1][j], table.field[i][j-1])
			table.field[i][j] = table.field[i][j] + min
		}
	}
	fmt.Println(table.field[x][y])

}
