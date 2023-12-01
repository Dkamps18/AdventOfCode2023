package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day1/input.txt")
	a := 0
	for _, v := range strings.Split(string(in), "\n") {
		d := []rune{}
		for _, c := range v {
			if c >= '0' && c <= '9' {
				d = append(d, c)
			}
		}
		dl := len(d)
		if dl == 1 {
			i, _ := strconv.Atoi(string(d) + string(d))
			a += i
		} else if dl == 2 {
			i, _ := strconv.Atoi(string(d))
			a += i
		} else if dl > 0 {
			i, _ := strconv.Atoi(string(d[0]) + string(d[dl-1]))
			a += i
		}
	}
	fmt.Println(a)
}
