package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day7/input.txt")
	// in, _ := os.ReadFile("./day7/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	cardstrength := map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 0, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

	var hands []hand
	for k, v := range lines {
		h := hand{ID: k + 1, CardsCount: make(map[rune]int)}
		ci := strings.Split(v, " ")
		for _, v := range ci[0] {
			h.Cards = append(h.Cards, v)
			h.CardsCount[v]++
		}
		if _, ok := h.CardsCount['J']; ok {
			var high rune
			for c, v := range h.CardsCount {
				if c != 'J' {
					if high == 0 {
						high = c
						continue
					}
					if v > h.CardsCount[high] {
						high = c
					}
				}
			}
			h.CardsCount[high] += h.CardsCount['J']
			delete(h.CardsCount, 'J')
		}
		switch len(h.CardsCount) {
		case 1:
			// Five of a kind
			h.Type = 6
			break
		case 2:
			for _, v := range h.CardsCount {
				if v == 4 {
					// Four of a kind
					h.Type = 5
					break
				}
			}
			if h.Type == 0 {
				// Full house
				h.Type = 4
			}
			break
		case 3:
			var high int
			for _, v := range h.CardsCount {
				if v > high {
					high = v
				}
			}
			switch high {
			case 3:
				// Three of a kind
				h.Type = 3
				break
			case 2:
				// Two pair
				h.Type = 2
				break
			}
			break
		case 4:
			// One pair
			h.Type = 1
		}
		b, _ := strconv.Atoi(ci[1])
		h.Bet = b
		hands = append(hands, h)
	}

	var rank []hand
	for _, v := range hands {
		rank = append(rank, v)
	}
	sort.Slice(rank, func(i, j int) bool {
		if rank[i].Type == rank[j].Type {
			for c, v := range rank[i].Cards {
				if v == rank[j].Cards[c] {
					continue
				}
				return cardstrength[v] < cardstrength[rank[j].Cards[c]]
			}
		}
		return rank[i].Type < rank[j].Type
	})
	a := 0
	for i := 1; i <= len(rank); i++ {
		a += i * rank[i-1].Bet
	}
	fmt.Println(a)
}

