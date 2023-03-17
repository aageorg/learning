package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gamer struct {
	cards []int
	name  string
}

func (g *Gamer) Show() int {
	first := -1
	if len(g.cards) > 0 {
		first = g.cards[0]
	}
	return first
}

func (g *Gamer) Take(a int, b int) {
	g.cards = append(g.cards[1:], []int{a, b}...)
}

func (g *Gamer) Lose() {
	g.cards = g.cards[1:]
}

func main() {
	A := Gamer{name: "first", cards: make([]int, 0, 10)}
	B := Gamer{name: "second", cards: make([]int, 0, 10)}

	reader := bufio.NewReader(os.Stdin)
	src, _, _ := reader.ReadLine()
	c1 := strings.Split(strings.TrimSpace(string(src)), " ")
	for _, c := range c1 {
		card, _ := strconv.Atoi(c)
		A.cards = append(A.cards, card)
	}
	src, _, _ = reader.ReadLine()
	c2 := strings.Split(strings.TrimSpace(string(src)), " ")
	for _, c := range c2 {
		card, _ := strconv.Atoi(c)
		B.cards = append(B.cards, card)
	}
	i := 0
	for i < 1000000 {
		fmt.Println(A.cards)
		fmt.Println(B.cards)
		a := A.Show()
		if a == -1 {
			fmt.Print(B.name + " ")
			fmt.Println(i)
			return
		}
		b := B.Show()
		if b == -1 {
			fmt.Print(A.name + " ")
			fmt.Println(i)
			return
		}
		if (a-b)*(a-b) == 81 {
			if a > b {
				A.Lose()
				B.Take(a, b)
			} else {
				B.Lose()
				A.Take(a, b)
			}
		} else {
			if a > b {
				B.Lose()
				A.Take(a, b)
			} else {
				A.Lose()
				B.Take(a, b)
			}
		}
		i++
	}
	fmt.Println("botva")
}
