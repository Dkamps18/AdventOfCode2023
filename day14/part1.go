package day14

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Part1() {
	// in, _ := os.ReadFile("./day14/input.txt")
	in, _ := os.ReadFile("./day14/ex.txt")
	rows := strings.Split(string(in), "\r\n")
	var cols []string
	for i := 0; i < len(rows[0]); i++ {
		str := ""
		for _, v := range rows {
			str += string(v[i])
		}
		cols = append(cols, str)
	}

	for k, v := range cols {
		var r []string
		parts := strings.Split(v, "#")
		for _, p := range parts {
			chars := strings.Split(p, "")
			sort.Sort(sort.Reverse(sort.StringSlice(chars)))
			r = append(r, strings.Join(chars, ""))
		}
		cols[k] = strings.Join(r, "#")
	}

	var a int
	for _, v := range cols {
		cl := len(v)
		for k, c := range v {
			if c == 'O' {
				a += cl - k
			}
		}
	}
	fmt.Println(a)
}
