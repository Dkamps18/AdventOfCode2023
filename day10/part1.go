package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day10/input.txt")
	// in, _ := os.ReadFile("./day10/ex3.txt")
	lines := strings.Split(string(in), "\r\n")
	grid := make(map[int][]rune)
	startx, starty := 0, 0
	for k, v := range lines {
		for k2, v := range v {
			grid[k] = append(grid[k], v)
			if v == 'S' {
				startx = k
				starty = k2
			}
		}
	}
	possteps := make(map[string]int)
	for _, v := range []rune{'n', 'e', 's', 'w'} {
		dir := v
		currx, curry := startx, starty
		maxx, maxy := len(lines), len(lines[0])
		var steps int
		going := true
		for going {
			switch dir {
			case 'n':
				if currx-1 < 0 {
					going = false
					break
				}
				np := grid[currx-1][curry]
				currx--
				switch np {
				case '|', 'S':
					break
				case 'F':
					dir = 'e'
					break
				case '7':
					dir = 'w'
					break
				default:
					going = false
					break
				}
				break
			case 'e':
				if curry+1 > maxy {
					going = false
					break
				}
				np := grid[currx][curry+1]
				curry++
				switch np {
				case '-', 'S':
					break
				case 'J':
					dir = 'n'
					break
				case '7':
					dir = 's'
					break
				default:
					going = false
					break
				}
				break
			case 's':
				if currx+1 > maxx {
					going = false
					break
				}
				np := grid[currx+1][curry]
				currx++
				switch np {
				case '|', 'S':
					break
				case 'L':
					dir = 'e'
					break
				case 'J':
					dir = 'w'
					break
				default:
					going = false
					break
				}
				break
			case 'w':
				if curry-1 < 0 {
					going = false
					break
				}
				np := grid[currx][curry-1]
				curry--
				switch np {
				case '-', 'S':
					break
				case 'F':
					dir = 's'
					break
				case 'L':
					dir = 'n'
					break
				default:
					going = false
					break
				}
				break
			}
			if !going || (currx == startx && curry == starty) {
				going = false
				break
			}
			steps++
			spos := strconv.Itoa(currx) + "-" + strconv.Itoa(curry)
			if val, ok := possteps[spos]; ok {
				if steps < val {
					possteps[spos] = steps
				}
			} else {
				possteps[spos] = steps
			}
		}
	}
	h := 0
	for _, v := range possteps {
		if v > h {
			h = v
		}
	}
	fmt.Println(h)
}
