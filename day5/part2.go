package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day5/input.txt")
	// in, _ := os.ReadFile("./day5/ex.txt")
	maps := strings.Split(string(in), "\r\n\r\n")
	var seedints []int
	var seedrange []locrange
	for _, v := range strings.Split(maps[0], " ")[1:] {
		s, _ := strconv.Atoi(v)
		seedints = append(seedints, s)
	}
	for i := 0; i < len(seedints); i++ {
		seedrange = append(seedrange, locrange{Dest: seedints[i], Src: seedints[i], RL: seedints[i+1]})
		i++
	}

	seedtosoil := getlocranges(maps[1])
	soiltofertilizer := getlocranges(maps[2])
	fertilizertowater := getlocranges(maps[3])
	watertolight := getlocranges(maps[4])
	lighttotemperature := getlocranges(maps[5])
	temperaturetohumidity := getlocranges(maps[6])
	humiditytolocation := getlocranges(maps[7])

	lowloc := 999999999999999999
	for _, v := range seedrange {
		for i := 0; i < v.RL; i++ {
			seed := v.Src + i
			soil := getval(seedtosoil, seed)
			fertilizer := getval(soiltofertilizer, soil)
			water := getval(fertilizertowater, fertilizer)
			light := getval(watertolight, water)
			temperature := getval(lighttotemperature, light)
			humidity := getval(temperaturetohumidity, temperature)
			location := getval(humiditytolocation, humidity)
			if location < lowloc {
				lowloc = location
			}
		}
	}
	fmt.Println(lowloc)
}

type locrange struct {
	Dest int
	Src  int
	RL   int
}

func getlocranges(m string) []locrange {
	r := []locrange{}
	for _, v := range strings.Split(m, "\r\n")[1:] {
		i := strings.Split(v, " ")
		dest, _ := strconv.Atoi(i[0])
		src, _ := strconv.Atoi(i[1])
		rl, _ := strconv.Atoi(i[2])
		r = append(r, locrange{Dest: dest, Src: src, RL: rl})
	}
	return r
}

func getval(r []locrange, i int) int {
	for _, v := range r {
		if i >= v.Src && i <= v.Src+v.RL {
			return v.Dest + (i - v.Src)
		}
	}
	return i
}
