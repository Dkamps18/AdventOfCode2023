package day14

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Part2() {
	in, _ := os.ReadFile("./day14/input.txt")
	// in, _ := os.ReadFile("./day14/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	grid := make(map[int][]rune)
	for k, v := range lines {
		grid[k] = []rune{}
		for _, c := range v {
			grid[k] = append(grid[k], c)
		}
	}
	var i int
	go func() {
		var l int
		for i < 1000000000 {
			time.Sleep(time.Second)
			fmt.Println(strconv.Itoa(i)+"/1000000000", 1000000000-i, "away", strconv.Itoa(i-l)+"/sec")
			l = i
		}
	}()
	rl := len(grid)
	for i = 0; i < 1000000000; i++ {
		for c := 0; c < len(grid[0]); c++ {
			parts := colparts(grid, c)
			var r [][]rune
			for _, p := range parts {
				c := p
				sort.Sort(sort.Reverse(RuneSlice(c)))
				r = append(r, c)
			}
			var ri int
			pl := len(parts)
			for k, v := range parts {
				for _, char := range v {
					grid[ri][c] = char
					ri++
				}
				if k+1 != pl {
					grid[ri][c] = '#'
					ri++
				}
			}
		}

		for r := 0; r < rl; r++ {
			parts := rowparts(grid[r])
			var r [][]rune
			for _, p := range parts {
				c := p
				sort.Sort(sort.Reverse(RuneSlice(c)))
				r = append(r, c)
			}
			var nr []rune
			pl := len(parts)
			for k, v := range r {
				nr = append(nr, v...)
				if k+1 != pl {
					nr = append(nr, '#')
				}
			}
		}

		for c := 0; c < len(grid[0]); c++ {
			parts := colparts(grid, c)
			var r [][]rune
			for _, p := range parts {
				c := p
				sort.Sort(RuneSlice(c))
				r = append(r, c)
			}
			var ri int
			pl := len(parts)
			for k, v := range parts {
				for _, char := range v {
					grid[ri][c] = char
					ri++
				}
				if k+1 != pl {
					grid[ri][c] = '#'
					ri++
				}
			}
		}

		for r := 0; r < rl; r++ {
			parts := rowparts(grid[r])
			var r [][]rune
			for _, p := range parts {
				c := p
				sort.Sort(RuneSlice(c))
				r = append(r, c)
			}
			var nr []rune
			pl := len(parts)
			for k, v := range r {
				nr = append(nr, v...)
				if k+1 != pl {
					nr = append(nr, '#')
				}
			}
		}
	}

	var a int
	cl := len(grid[0])
	for c := 0; c < cl; c++ {
		for r := 0; r < rl; r++ {
			if grid[r][c] == 'O' {
				a += rl - r
			}
		}
	}
	fmt.Println(a)

}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func colparts(grid map[int][]rune, c int) [][]rune {
	var parts [][]rune
	var t []rune
	for r := 0; r < len(grid); r++ {
		if grid[r][c] != '#' {
			t = append(t, grid[r][c])
		} else {
			parts = append(parts, t)
			t = []rune{}
		}
	}
	if len(t) != 0 {
		parts = append(parts, t)
	}
	return parts
}

func rowparts(row []rune) [][]rune {
	var parts [][]rune
	var t []rune
	for _, v := range row {
		if v != '#' {
			t = append(t, v)
		} else {
			parts = append(parts, t)
			t = []rune{}
		}
	}
	if len(t) != 0 {
		parts = append(parts, t)
	}
	return parts
}
