package day2

import (
	"adventofcode2022/internal/util"
	"fmt"
)

// func getScore(opp string, you string) int {

// }
const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
	Win          = 6
	Lose         = 0
	Draw         = 3
)

func Solve() {

	lines := util.ReadLines("2")
	m := map[string]int{
		"A X": Rock + Draw,
		"A Y": Paper + Win,
		"A Z": Scissors + Lose,
		"B X": Rock + Lose,
		"B Y": Paper + Draw,
		"B Z": Scissors + Win,
		"C X": Rock + Win,
		"C Y": Paper + Lose,
		"C Z": Scissors + Draw,
	}
	val := 0
	for _, line := range lines {
		val += m[line]
	}
	fmt.Println("Score:", val)
}
