package main

import (
	"adventofcode2022/days/day1"
	"adventofcode2022/days/day2"
	"adventofcode2022/days/day3"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	fmt.Println("Solving for Day ", day)
	switch day {
	case "1":
		day1.Solve()
	case "2":
		day2.Solve()
	case "3":
		day3.Solve()
	default:
		panic("Invalid Day Entered")
	}
}
