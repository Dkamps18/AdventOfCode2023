package main

import (
	"Dkamps18/AdventOfCode2023/day1"
	"os"
)

func main() {
	switch os.Args[1] {
	case "day1-1":
		day1.Part1()
		break
	case "day1-2":
		day1.Part2()
		break
	}
}
