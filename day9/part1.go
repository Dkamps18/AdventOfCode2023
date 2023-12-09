package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day9/input.txt")
	// in, _ := os.ReadFile("./day9/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	var seqs []seq
	for _, v := range lines {
		s := seq{Lines: make(map[int][]int)}
		s.Lines[0] = []int{}
		for _, ds := range strings.Split(v, " ") {
			d, _ := strconv.Atoi(ds)
			s.Lines[0] = append(s.Lines[0], d)
		}
		seqs = append(seqs, s)
	}
	a := 0
	for _, s := range seqs {
		zeroes := false
		i := 0
		for !zeroes {
			s.Lines[i+1] = []int{}
			ll := len(s.Lines[i])
			for k, v := range s.Lines[i] {
				if k == ll-1 {
					break
				}
				s.Lines[i+1] = append(s.Lines[i+1], s.Lines[i][k+1]-v)
			}
			nz := false
			for _, v := range s.Lines[i+1] {
				if v != 0 {
					nz = true
					break
				}
			}
			if nz {
				i++
			} else {
				zeroes = true
			}
		}
		for i := len(s.Lines) - 1; i > 0; i-- {
			s.Lines[i-1] = append(s.Lines[i-1], s.Lines[i][len(s.Lines[i])-1]+s.Lines[i-1][len(s.Lines[i-1])-1])
		}
		a += s.Lines[0][len(s.Lines[0])-1]
	}
	fmt.Println(a)
}

type seq struct {
	Lines map[int][]int
}
