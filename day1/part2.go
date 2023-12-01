package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day1/input.txt")
	a := 0
	for _, v := range strings.Split(string(in), "\n") {
		d := []byte{}
		vl := len(v)
		for i := 0; i < vl; i++ {
			switch v[i] {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				d = append(d, v[i])
				break
			case 'o':
				if vl-i >= 3 {
					if v[i:i+3] == "one" {
						d = append(d, '1')
					}
				}
				break
			case 't':
				if vl-i >= 5 {
					if v[i:i+5] == "three" {
						d = append(d, '3')
					} else if v[i:i+3] == "two" {
						d = append(d, '2')
					}
				} else if vl-i >= 3 {
					if v[i:i+3] == "two" {
						d = append(d, '2')
					}
				}
				break
			case 'f':
				if vl-i >= 4 {
					if v[i:i+4] == "four" {
						d = append(d, '4')
					} else if v[i:i+4] == "five" {
						d = append(d, '5')
					}
				}
				break
			case 's':
				if vl-i >= 5 {
					if v[i:i+5] == "seven" {
						d = append(d, '7')
					} else if v[i:i+3] == "six" {
						d = append(d, '6')
					}
				} else if vl-i >= 3 {
					if v[i:i+3] == "six" {
						d = append(d, '6')
					}
				}
				break
			case 'e':
				if vl-i >= 5 {
					if v[i:i+5] == "eight" {
						d = append(d, '8')
					}
				}
				break
			case 'n':
				if vl-i >= 4 {
					if v[i:i+4] == "nine" {
						d = append(d, '9')
					}
				}
				break
			}
		}
		// break
		dl := len(d)
		if dl == 1 {
			i, _ := strconv.Atoi(string(d) + string(d))
			a += i
		} else if dl == 2 {
			i, _ := strconv.Atoi(string(d))
			a += i
		} else if dl > 2 {
			i, _ := strconv.Atoi(string(d[0]) + string(d[dl-1]))
			a += i
		}
	}
	fmt.Println(a)
}
