package day1

import (
	"adventofcode2022/internal/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	scanner := util.GetInput("1")

	scanner.Split(bufio.ScanLines)

	max_calories := -1
	cur := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			if cur > max_calories {
				max_calories = cur
			}

			cur = 0
		} else {
			iv, err := strconv.Atoi(line)
			util.Check(err)

			cur = cur + iv
		}
	}
	fmt.Println("Max Calories:", max_calories)
}
