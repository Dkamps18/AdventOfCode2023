package day12

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	// in, _ := os.ReadFile("./day12/input.txt")
	in, _ := os.ReadFile("./day12/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	a := make(map[int]int)
	for k, v := range lines {
		r := strings.Split(v, " ")
		var rd []int
		for _, v := range strings.Split(r[1], ",") {
			vd, _ := strconv.Atoi(v)
			rd = append(rd, vd)
		}
		a[k] += process(r[0], rd)
	}
	var t int
	for _, v := range a {
		t += int(math.Pow(float64(v), 5))
	}
	fmt.Println(t)
}

func process(i string, d []int) int {
	if i == "" {
		if len(d) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(d) == 0 {
		if strings.Contains(i, "#") {
			return 0
		} else {
			return 1
		}
	}

	var r int
	switch i[0] {
	case '.', '?':
		r += process(i[1:], d)
		break
	}
	switch i[0] {
	case '#', '?':
		if d[0] <= len(i) && (d[0] == len(i) || i[d[0]] != '#') && !strings.Contains(i[:d[0]], ".") {
			if d[0] == len(i) {
				r += process("", d[1:])
			} else {
				r += process(i[d[0]+1:], d[1:])
			}
		}
		break
	}
	return r
}
