package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day4/input.txt")
	// in, _ := os.ReadFile("./day4/ex.txt")
	lines := strings.Split(strings.ReplaceAll(string(in), "  ", " "), "\r\n")
	cards := make(map[int]int)
	for k, v := range lines {
		cn := k + 1
		cards[cn]++
		ci := strings.SplitN(v, " ", 2)
		cv := strings.SplitN(ci[1], " | ", 2)
		var winning []int
		var numbers []int
		for _, n := range strings.Split(cv[0], " ") {
			cn, _ := strconv.Atoi(n)
			winning = append(winning, cn)
		}
		for _, n := range strings.Split(cv[1], " ") {
			cn, _ := strconv.Atoi(n)
			numbers = append(numbers, cn)
		}

		matches := 0
		for _, n := range numbers {
			if inarray(n, winning) {
				matches++
			}
		}
		fmt.Println(cn, matches)

		if matches != 0 {
			for c := 0; c < cards[cn]; c++ {
				for i := 1; i <= matches; i++ {
					cards[cn+i]++
				}
			}
		}

	}
	fmt.Println(cards)
	totalcards := 0
	for _, v := range cards {
		totalcards += v
	}
	fmt.Println(totalcards)
}
