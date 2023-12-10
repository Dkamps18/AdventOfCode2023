package main

import (
	"Dkamps18/AdventOfCode2023/day1"
	"Dkamps18/AdventOfCode2023/day10"
	"Dkamps18/AdventOfCode2023/day2"
	"Dkamps18/AdventOfCode2023/day3"
	"Dkamps18/AdventOfCode2023/day4"
	"Dkamps18/AdventOfCode2023/day5"
	"Dkamps18/AdventOfCode2023/day6"
	"Dkamps18/AdventOfCode2023/day7"
	"Dkamps18/AdventOfCode2023/day8"
	"Dkamps18/AdventOfCode2023/day9"
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
	case "day3-1":
		day3.Part1()
		break
	case "day3-2":
		day3.Part2()
		break
	case "day4-1":
		day4.Part1()
		break
	case "day4-2":
		day4.Part2()
		break
	case "day5-1":
		day5.Part1()
		break
	case "day5-2":
		day5.Part2()
		break
	case "day6-1":
		day6.Part1()
		break
	case "day6-2":
		day6.Part2()
		break
	case "day7-1":
		day7.Part1()
		break
	case "day7-2":
		day7.Part2()
		break
	case "day8-1":
		day8.Part1()
		break
	case "day8-2":
		day8.Part2()
		break
	case "day9-1":
		day9.Part1()
		break
	case "day9-2":
		day9.Part2()
		break
	case "day10-1":
		day10.Part1()
		break
	case "day10-2":
		day10.Part2()
		break
	default:
		fmt.Println("Not found")
		os.Exit(1)
	}
}
