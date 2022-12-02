package main

import (
	"adventofcode2022/days/day1"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	fmt.Println("Solving for Day ", day)
	switch day {
	case "1":
		day1.Solve()
	}
}
