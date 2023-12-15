package day15

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day15/input.txt")
	// in, _ := os.ReadFile("./day15/ex.txt")
	groups := strings.Split(string(in), ",")

	b := make([][]string, 256)
	fl := make(map[string]int)

	for _, g := range groups {
		if strings.Contains(g, "-") {
			l := g[:len(g)-1]
			i := hash(l)
			for k, v := range b[i] {
				if l == v {
					b[i] = append(b[i][:k], b[i][k+1:]...)
					break
				}
			}
		} else {
			p := strings.Split(g, "=")
			i := hash(p[0])
			if !inarray(p[0], b[i]) {
				b[i] = append(b[i], p[0])
			}
			l, _ := strconv.Atoi(p[1])
			fl[p[0]] = l
		}
	}

	a := 0
	for k, v := range b {
		for k2, v2 := range v {
			a += (k + 1) * (k2 + 1) * fl[v2]
		}
	}
	fmt.Println(a)
}

func hash(v string) int {
	var t int
	for _, c := range v {
		t += int(c)
		t = t * 17
		t = t % 256
	}
	return t
}

func inarray(s string, a []string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}
	return false
}
