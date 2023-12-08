package day8

import (
	"fmt"
	"os"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day8/input.txt")
	// in, _ := os.ReadFile("./day8/ex.txt")
	parts := strings.Split(string(in), "\r\n\r\n")
	// fmt.Println(parts)
	lrs := strings.Split(parts[0], "")
	locs := make(map[string]loc)
	for _, v := range strings.Split(parts[1], "\r\n") {
		li := v[0:3]
		locs[li] = loc{Left: v[7:10], Right: v[12:15]}
	}

	found := false
	at := "AAA"
	steps := 0
	for !found {
		for _, v := range lrs {
			switch v {
			case "L":
				at = locs[at].Left
				break
			case "R":
				at = locs[at].Right
				break
			}
			steps++
			if at == "ZZZ" {
				found = true
				break
			}
		}
	}
	fmt.Println(steps)
}

type loc struct {
	Left  string
	Right string
}
