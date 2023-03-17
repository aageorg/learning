package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	schema := make(map[int]int)
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, b := range input {
		if b == 10 || b == 32 {
			continue
		}
		if _, ok := schema[int(b)]; !ok {
			schema[int(b)] = 0
		}
		schema[int(b)]++
	}
	var output string
	var keys []int

	for k, _ := range schema {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		output += string(k)
	}
External:
	for {
		var str string
		for _, symb := range keys {
			if schema[symb] > 0 {
				str += "#"
				schema[symb]--
			} else {
				str += " "
			}
		}
		if strings.TrimSpace(str) != "" {
			output = str + "\n" + output
		} else {

			break External
		}
	}
	fmt.Println(output)
}
