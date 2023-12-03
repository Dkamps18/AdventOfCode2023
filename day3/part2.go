package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	in, _ := os.ReadFile("./day3/input.txt")
	// in, _ := os.ReadFile("./day3/ex.txt")
	lines := strings.Split(string(in), "\r\n")
	ll := len(lines)
	astrixes := []pos{}
	digits := make(map[int]map[int]int)
	for l, v := range lines {
		digits[l] = make(map[int]int)
		vl := len(v)
		for i := 0; i < vl; i++ {
			switch v[i] {
			case '*':
				astrixes = append(astrixes, pos{Line: l, Index: i})
				break
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var npos []int
				npos = append(npos, i)
				fn := true
				o := 1
				for fn {
					if v[i+o] >= '0' && v[i+o] <= '9' {
						npos = append(npos, i+o)
						o++
					} else {
						fn = false
					}
					if i+o == vl {
						fn = false
					}
				}
				nposl := len(npos)
				pn, _ := strconv.Atoi(v[npos[0] : npos[nposl-1]+1])
				for _, p := range npos {
					digits[l][p] = pn
				}
				i += nposl - 1
			}
		}
	}

	var a int
	gc := 0
	for _, p := range astrixes {
		vl := len(lines[p.Line])
		var tl, t, tr, l, r, bl, b, br int
		if p.Line != 0 {
			if p.Index-1 >= 0 {
				if val, ok := digits[p.Line-1][p.Index-1]; ok {
					tl = val
				}
			}
			if val, ok := digits[p.Line-1][p.Index]; ok {
				t = val
			}
			if p.Index+1 <= vl {
				if val, ok := digits[p.Line-1][p.Index+1]; ok {
					tr = val
				}
			}
		}
		if p.Index-1 >= 0 {
			if val, ok := digits[p.Line][p.Index-1]; ok {
				l = val
			}
		}
		if p.Index+1 <= vl {
			if val, ok := digits[p.Line][p.Index+1]; ok {
				r = val
			}
		}
		if p.Line+1 < ll {
			if p.Index-1 >= 0 {
				if val, ok := digits[p.Line+1][p.Index-1]; ok {
					bl = val
				}
			}
			if val, ok := digits[p.Line+1][p.Index]; ok {
				b = val
			}
			if p.Index+1 <= vl {
				if val, ok := digits[p.Line+1][p.Index+1]; ok {
					br = val
				}
			}
		}

		da := []int{}
		if t > 0 {
			da = append(da, t)
		} else {
			if tl > 0 {
				da = append(da, tl)
			}
			if tr > 0 {
				da = append(da, tr)
			}
		}
		if l > 0 {
			da = append(da, l)
		}
		if r > 0 {
			da = append(da, r)
		}
		if b > 0 {
			da = append(da, b)
		} else {
			if bl > 0 {
				da = append(da, bl)
			}
			if br > 0 {
				da = append(da, br)
			}
		}

		if len(da) == 2 {
			gc++
			a += da[0] * da[1]
		}
	}
	fmt.Println(gc)
	fmt.Println(a)
}

type pos struct {
	Line  int
	Index int
}

// not
// 60387533 (low)
// 58944830 (low)
// 124035795

// 66769383
