package day3

import (
	"adventofcode2022/internal/util"
	"fmt"
	"sort"
	"unicode"

	"golang.org/x/exp/slices"
)

func initSlice(size int) []int {
	agg := make([]int, size)
	for i := 0; i < len(agg); i++ {
		agg[i] = 0
	}
	return agg
}
func getPriority(c rune) int {
	var cc rune
	if unicode.IsUpper(c) {
		cc = unicode.ToLower(c)
	} else {
		cc = unicode.ToUpper(c)
	}
	var shift int
	if cc > 90 {
		shift = 70
	} else {
		shift = 64
	}
	return int(cc) - shift
}
func rune2string(c []rune) []string {
	var ret []string
	for _, v := range c {
		ret = append(ret, string(v))
	}
	return ret
}
func calculate(c []rune) int {
	ret := 0
	mid := len(c) / 2
	c1 := c[0:mid]
	c2 := c[mid:]
	var visited []rune
	for _, v := range c1 {
		if !slices.Contains(visited, v) {
			visited = append(visited, v)
			if slices.Contains(c2, v) {
				ret += getPriority(v)
			}
		}
	}
	fmt.Println(ret, rune2string(c1), rune2string(c2))

	return ret
}
func part1(lines []string) {
	tot := 0
	for _, line := range lines {
		cs := []rune(line)
		priority := calculate(cs)
		tot += priority
	}
	fmt.Println("Total:", tot)
}

type sortedRunes []rune

func (a sortedRunes) Len() int {
	return len(a)
}

func (a sortedRunes) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a sortedRunes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func processGroup(group []string) rune {
	rg := prepareGroup(group)

	for _, v := range rg[0] {
		if slices.Contains(rg[1], v) && slices.Contains(rg[2], v) {
			return v
		}
	}
	return rune(0)
}

func prepareGroup(group []string) [][]rune {
	var ret [][]rune
	for _, v := range group {
		g := []rune(v)
		sort.Sort(sortedRunes(g))
		g = slices.Compact(g)
		ret = append(ret, g)
	}
	return ret
}

func part2(lines []string) {
	groupSize := 3
	tot := 0
	for i := 0; i < len(lines); i += groupSize {
		r := processGroup(lines[i : i+groupSize])
		tot += getPriority(r)
		fmt.Println(i, string(r))
	}
	fmt.Println("Part2:", tot)
}

func Solve() {
	lines := util.ReadLines("3")
	//part1(lines)
	part2(lines)
}
