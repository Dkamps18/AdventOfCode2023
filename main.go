package main

import (
	"Dkamps18/AdventOfCode2023/day1"
	"Dkamps18/AdventOfCode2023/day2"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify as dayx-x")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "day1-1":
		day1.Part1()
		break
	case "day1-2":
		day1.Part2()
		break
	case "day2-1":
		day2.Part1()
		break
	case "day2-2":
		day2.Part2()
		break
	default:
		fmt.Println("Not found")
		os.Exit(1)
	}
}
