package day11

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day11/input.txt")
	// in, _ := os.ReadFile("./day11/ex.txt")
	grid := strings.Split(string(in), "\r\n")
	var er []int
	var ec []int
	for k, v := range grid {
		if !strings.Contains(v, "#") {
			er = append(er, k)
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		var g bool
		for _, v := range grid {
			if v[i] == '#' {
				g = true
			}
		}
		if !g {
			ec = append(ec, i)
		}
	}

	var galaxies []pos

	for x, v := range grid {
		for y, c := range v {
			if c == '#' {
				erc := 0
				for _, v := range er {
					if v < x {
						erc++
					}
				}
				ecc := 0
				for _, v := range ec {
					if v < y {
						ecc++
					}
				}
				galaxies = append(galaxies, pos{X: x + (erc * 999999), Y: y + (ecc * 999999)})
			}
		}
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
