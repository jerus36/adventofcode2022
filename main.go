package main

import (
	"adventofcode2022/days/day1"
	"adventofcode2022/days/day2"
	"adventofcode2022/days/day3"
	"adventofcode2022/days/day4"
	"adventofcode2022/internal/util"
	"fmt"
	"os"
)

var days = map[string]interface{}{
	"1": day1.Solve,
	"2": day2.Solve,
	"3": day3.Solve,
	"4": day4.Solve,
}

func notImplemented(day string) {
	s := fmt.Sprintf("Day %s not implmented yet", day)
	panic(s)
}

func main() {
	day := os.Args[1]
	if f, ok := days[day]; ok {
		fmt.Println("Solving for Day ", day)
		lines := util.ReadLines(day)
		f.(func([]string))(lines)
	} else {
		notImplemented(day)
	}
}
