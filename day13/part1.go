package day13

import (
	"fmt"
	"os"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day13/input.txt")
	// in, _ := os.ReadFile("./day13/ex2.txt")
	var rows []int
	var cols []int
	for _, pattern := range strings.Split(string(in), "\r\n\r\n") {
		r := strings.Split(pattern, "\r\n")
		var c []string
		for i := 0; i < len(r[0]); i++ {
			str := ""
			for _, v := range r {
				str += string(v[i])
			}
			c = append(c, str)
		}

		rl := len(r)
		cl := len(c)

		for k, v := range r {
			if k-1 >= 0 && r[k-1] == v {
				i := 1
				m := true
				for m && k-i-1 >= 0 && k+i < rl {
					m = r[k-i-1] == r[k+i]
					i++
				}
				if m {
					rows = append(rows, k)
					break
				}
			}
		}

		for k, v := range c {
			if k-1 >= 0 && c[k-1] == v {
				i := 1
				m := true
				for m && k-i-1 >= 0 && k+i < cl {
					m = c[k-i-1] == c[k+i]
					i++
				}
				if m {
					cols = append(cols, k)
					break
				}
			}
		}
	}

	var rc, cc int
	for _, v := range rows {
		rc += v
	}
	for _, v := range cols {
		cc += v
	}

	fmt.Println(cc + 100*rc)
}
