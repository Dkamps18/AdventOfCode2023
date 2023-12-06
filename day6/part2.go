package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day6/input.txt")
	// in, _ := os.ReadFile("./day6/ex.txt")
	lines := strings.Split(strings.ReplaceAll(string(in), " ", ""), "\r\n")
	time, _ := strconv.Atoi(strings.Split(lines[0], ":")[1])
	distance, _ := strconv.Atoi(strings.Split(lines[1], ":")[1])

	posibilities := 0
	for i := 1; i < time; i++ {
		if (time-i)*i > distance {
			posibilities++
		}
	}
	fmt.Println(posibilities)
}
