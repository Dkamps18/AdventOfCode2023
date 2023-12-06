package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day6/input.txt")
	// in, _ := os.ReadFile("./day6/ex.txt")
	// For some stupid reason doesn't work like it did on day 4
	// lines := strings.Split(strings.ReplaceAll(string(in), "  ", " "), "\r\n")
	// times := strings.Split(lines[0], " ")
	// distances := strings.Split(lines[1], " ")
	lines := strings.Split(string(in), "\r\n")
	var times []string
	for _, v := range strings.Split(lines[0], " ") {
		if v != "" {
			times = append(times, v)
		}
	}
	var distances []string
	for _, v := range strings.Split(lines[1], " ") {
		if v != "" {
			distances = append(distances, v)
		}
	}

	var races []race

	for k, v := range times[1:] {
		t, _ := strconv.Atoi(v)
		d, _ := strconv.Atoi(distances[k+1])
		races = append(races, race{Time: t, Distance: d})
	}

	posibilities := make(map[int]int)
	for k, v := range races {
		for i := 1; i < v.Time; i++ {
			if (v.Time-i)*i > v.Distance {
				posibilities[k]++
			}
		}
	}

	a := posibilities[0]
	for i := 1; i < len(posibilities); i++ {
		a = a * posibilities[i]
	}
	fmt.Println(a)
}

type race struct {
	Time     int
	Distance int
}
