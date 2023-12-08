package day8

import (
	"fmt"
	"os"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day8/input.txt")
	// in, _ := os.ReadFile("./day8/ex2.txt")
	parts := strings.Split(string(in), "\r\n\r\n")
	// fmt.Println(parts)
	lrs := strings.Split(parts[0], "")
	locs := make(map[string]loc)
	starts := 0
	at := make(map[int]string)
	for _, v := range strings.Split(parts[1], "\r\n") {
		li := v[0:3]
		if li[2] == 'A' {
			at[starts] = li
			starts++
		}
		locs[li] = loc{Left: v[7:10], Right: v[12:15]}
	}

	steps := make(map[int]int)
	for i := 0; i < starts; i++ {
		found := false
		for !found {
			for _, v := range lrs {
				switch v {
				case "L":
					at[i] = locs[at[i]].Left
					break
				case "R":
					at[i] = locs[at[i]].Right
					break
				}
				steps[i]++
				if at[i][2] == 'Z' {
					found = true
				}
			}
		}
	}
	var minsteps []int
	for _, v := range steps {
		minsteps = append(minsteps, v)
	}
	fmt.Println(LCM(minsteps[0], minsteps[1], minsteps[2:]...))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
