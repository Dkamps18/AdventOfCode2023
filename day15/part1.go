package day15

import (
	"fmt"
	"os"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day15/input.txt")
	// in, _ := os.ReadFile("./day15/ex.txt")
	groups := strings.Split(string(in), ",")
	var a int
	for _, v := range groups {
		var t int
		for _, c := range v {
			t += int(c)
			t = t * 17
			t = t % 256
		}
		a += t
	}
	fmt.Println(a)
}
