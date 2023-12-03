package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day3/input.txt")
	// in, _ := os.ReadFile("./day3/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	ln := len(lines)
	var parts []int
	for l, v := range lines {
		vl := len(v)
		for i := 0; i < vl; i++ {
			if v[i] >= '0' && v[i] <= '9' {
				var npos []int
				npos = append(npos, i)
				fn := true
				o := 1
				for fn {
					if v[i+o] >= '0' && v[i+o] <= '9' {
						npos = append(npos, i+o)
						o++
					} else {
						fn = false
					}
					if i+o == vl {
						fn = false
					}
				}
				nposl := len(npos)
				if ispart(lines, ln, l, npos, nposl, vl-1) {
					pn, _ := strconv.Atoi(v[npos[0] : npos[nposl-1]+1])
					parts = append(parts, pn)
				}
				i += nposl - 1
			}
		}
	}
	a := 0
	for _, v := range parts {
		a += v
	}
	fmt.Println(a)
}

func ispart(lines []string, ln int, l int, npos []int, nposl int, ll int) bool {
	if l != 0 {
		if npos[0] != 0 {
			if issymbol(lines[l-1][npos[0]-1]) {
				return true
			}
		}
		for _, p := range npos {
			if issymbol(lines[l-1][p]) {
				return true
			}
		}
		if npos[nposl-1] != ll {
			if issymbol(lines[l-1][npos[nposl-1]+1]) {
				return true
			}
		}
	}
	if npos[0] != 0 {
		if issymbol(lines[l][npos[0]-1]) {
			return true
		}
	}
	if npos[nposl-1] != ll {
		if issymbol(lines[l][npos[nposl-1]+1]) {
			return true
		}
	}

	if l != ln-1 {
		if npos[0] != 0 {
			if issymbol(lines[l+1][npos[0]-1]) {
				return true
			}
		}
		for _, p := range npos {
			if issymbol(lines[l+1][p]) {
				return true
			}
		}

		if npos[nposl-1] != ll {
			if issymbol(lines[l+1][npos[nposl-1]+1]) {
				return true
			}
		}
	}
	return false
}

func issymbol(s byte) bool {
	return !(s >= '0' && s <= '9') && s != '.'
}
