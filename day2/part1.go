package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	var games []game
	in, _ := os.ReadFile("./day2/input.txt")
	// in, _ := os.ReadFile("./day2/ex.txt")
	for _, v := range strings.Split(string(in), "\r\n") {
		g := game{}
		gi := strings.SplitN(v, " ", 3)
		id, _ := strconv.Atoi(gi[1][:len(gi[1])-1])
		g.ID = id
		for _, h := range strings.Split(gi[2], "; ") {
			for _, cs := range strings.Split(h, ", ") {
				// fmt.Println(cs)
				ci := strings.SplitN(cs, " ", 2)
				a, _ := strconv.Atoi(ci[0])
				switch ci[1] {
				case "red":
					if a > g.Red {
						g.Red = a
					}
					break
				case "green":
					if a > g.Green {
						g.Green = a
					}
					break
				case "blue":
					if a > g.Blue {
						g.Blue = a
					}
					break
				}
			}
		}
		games = append(games, g)
	}

	possible := []int{}
	for _, v := range games {
		if v.Red <= 12 && v.Green <= 13 && v.Blue <= 14 {
			possible = append(possible, v.ID)
		}
	}
	a := 0
	for _, v := range possible {
		a += v
	}
	fmt.Println(a)
}

type game struct {
	ID    int
	Red   int
	Green int
	Blue  int
}
