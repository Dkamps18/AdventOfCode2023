package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day12/input.txt")
	// in, _ := os.ReadFile("./day12/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	a := make(map[int]int)
	for k, v := range lines {
		ir := strings.Split(v, " ")
		r := make([]string, 2)
		var tir0 []string
		var tir1 []string
		for i := 0; i < 5; i++ {
			tir0 = append(tir0, ir[0])
			tir1 = append(tir1, ir[1])
		}
		r[0] = strings.Join(tir0, "?")
		r[1] = strings.Join(tir1, ",")
		var rd []int
		for _, v := range strings.Split(r[1], ",") {
			vd, _ := strconv.Atoi(v)
			rd = append(rd, vd)
		}
		a[k] += process(r[0], rd)
	}
	var t int
	for _, v := range a {
		t += v
	}
	fmt.Println(t)
}
