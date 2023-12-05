package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	in, _ := os.ReadFile("./day5/input.txt")
	// in, _ := os.ReadFile("./day5/ex.txt")
	maps := strings.Split(string(in), "\r\n\r\n")
	seedints := make(map[int]int)
	for k, v := range strings.Split(maps[0], " ")[1:] {
		s, _ := strconv.Atoi(v)
		seedints[k] = s
	}
	seedtosoil := mapinput(maps[1], seedints)
	soiltofertilizer := mapinput(maps[2], seedtosoil)
	fertilizertowater := mapinput(maps[3], soiltofertilizer)
	watertolight := mapinput(maps[4], fertilizertowater)
	lighttotemperature := mapinput(maps[5], watertolight)
	temperaturetohumidity := mapinput(maps[6], lighttotemperature)
	humiditytolocation := mapinput(maps[7], temperaturetohumidity)

	var seeds []seed

	for _, v := range seedints {
		s := seed{}
		s.Seed = v
		if val, ok := seedtosoil[s.Seed]; ok {
			s.Soil = val
		} else {
			s.Soil = s.Seed
		}
		if val, ok := soiltofertilizer[s.Soil]; ok {
			s.Fertilizer = val
		} else {
			s.Fertilizer = s.Soil
		}
		if val, ok := fertilizertowater[s.Fertilizer]; ok {
			s.Water = val
		} else {
			s.Water = s.Fertilizer
		}
		if val, ok := watertolight[s.Water]; ok {
			s.Light = val
		} else {
			s.Light = s.Water
		}
		if val, ok := lighttotemperature[s.Light]; ok {
			s.Temperature = val
		} else {
			s.Temperature = s.Light
		}
		if val, ok := temperaturetohumidity[s.Temperature]; ok {
			s.Humidity = val
		} else {
			s.Humidity = s.Temperature
		}
		if val, ok := humiditytolocation[s.Humidity]; ok {
			s.Location = val
		} else {
			s.Location = s.Humidity
		}
		seeds = append(seeds, s)
	}

	lowloc := 999999999999999999
	for _, v := range seeds {
		if v.Location < lowloc {
			lowloc = v.Location
		}
	}
	fmt.Println(lowloc)
}

type seed struct {
	Seed        int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

func mapinput(m string, in map[int]int) map[int]int {
	r := make(map[int]int)
	sd := make(map[int]int)
	srl := make(map[int]int)
	for _, v := range strings.Split(m, "\r\n")[1:] {
		i := strings.Split(v, " ")
		dest, _ := strconv.Atoi(i[0])
		src, _ := strconv.Atoi(i[1])
		rl, _ := strconv.Atoi(i[2])
		sd[src] = dest
		srl[src] = rl
	}
	for _, v := range in {
		for s, d := range sd {
			if v >= s && v <= s+srl[s] {
				r[v] = d + (v - s)
				break
			}
		}
		if _, ok := r[v]; !ok {
			r[v] = v
		}
	}
	return r
}
