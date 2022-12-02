package day2

import (
	"adventofcode2022/internal/util"
	"fmt"
	"sort"
)

const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
	Win          = 0
	Lose         = 1
)
const (
	A string = "Rock"
	B        = "Paper"
	C        = "Scissors"
)

// func getScore(opp string, you string) int {

// }

func Solve() {
	lines := util.ReadLines("2")
	sort.Strings(lines)

	for _, line := range lines {
		fmt.Println(line)
	}
}
