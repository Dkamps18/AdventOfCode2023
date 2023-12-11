package day11

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day11/input.txt")
	// in, _ := os.ReadFile("./day11/ex.txt")
	grid := strings.Split(string(in), "\r\n")
	var t []string
	var t2 [][]byte
	for _, v := range grid {
		if strings.Contains(v, "#") {
			t2 = append(t2, []byte{})
			t = append(t, v)
		} else {
			t2 = append(t2, []byte{}, []byte{})
			t = append(t, v, v)
		}
	}
	grid = t
	for i := 0; i < len(grid[0]); i++ {
		var g bool
		for k, v := range grid {
			if v[i] == '#' {
				g = true
			}
			t2[k] = append(t2[k], v[i])
		}
		if !g {
			for k := range grid {
				t2[k] = append(t2[k], '.')
			}
		}
	}

	var galaxies []pos

	for k, v := range t2 {
		for y, c := range v {
			if c == '#' {
				galaxies = append(galaxies, pos{k + 1, y + 1})
			}
		}
		grid[k] = string(v)
	}

	var a int
	galaxydistances := make(map[string]int)
	for g := range galaxies {
		for g2 := range galaxies {
			if g == g2 {
				continue
			}
			var gk string
			if g > g2 {
				gk = strconv.Itoa(g) + "-" + strconv.Itoa(g2)
			} else {
				gk = strconv.Itoa(g2) + "-" + strconv.Itoa(g)
			}
			if _, ok := galaxydistances[gk]; !ok {
				galaxydistances[gk] = int(math.Abs(float64(galaxies[g2].X-galaxies[g].X) + math.Abs(float64(galaxies[g2].Y-galaxies[g].Y))))
			}
		}
	}
	for _, v := range galaxydistances {
		a += v
	}
	fmt.Println(a)
}

type pos struct {
	X int
	Y int
}
