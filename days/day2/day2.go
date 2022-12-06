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

func part2_map() map[string]int {
	m := map[string]int{
		"A X": Scissors + Lose,
		"A Y": Rock + Draw,
		"A Z": Paper + Win,
		"B X": Rock + Lose,
		"B Y": Paper + Draw,
		"B Z": Scissors + Win,
		"C X": Paper + Lose,
		"C Y": Scissors + Draw,
		"C Z": Rock + Win,
	}
	return m
}

func part1_map() map[string]int {
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
	return m
}

func Solve(l []string) {
	lines := util.ReadLines("2")
	m := part1_map()
	val := 0
	for _, line := range lines {
		val += m[line]
	}
	fmt.Println("Part1:", val)

	m = part2_map()
	val = 0
	for _, line := range lines {
		val += m[line]
	}
	fmt.Println("Part2:", val)

}
