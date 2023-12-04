package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day4/input.txt")
	// in, _ := os.ReadFile("./day4/ex.txt")
	lines := strings.Split(strings.ReplaceAll(string(in), "  ", " "), "\r\n")

	var totalpoints int
	for _, v := range lines {
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
		points := 0
		for _, n := range numbers {
			if inarray(n, winning) {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}
		totalpoints += points
	}
	fmt.Println(totalpoints)
}

func inarray(n int, a []int) bool {
	for _, v := range a {
		if n == v {
			return true
		}
	}
	return false
}
