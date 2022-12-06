package day1

import (
	"adventofcode2022/internal/util"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve(lines []string) {
	scanner := util.GetInput("1")

	scanner.Split(bufio.ScanLines)

	cur := 0
	var res []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			res = append(res, cur)
			cur = 0
		} else {
			iv, err := strconv.Atoi(line)
			util.Check(err)

			cur = cur + iv
		}
	}
	sort.Ints(res)
	tot := 0
	for _, v := range res[len(res)-3:] {
		tot += v
	}
	fmt.Println("Part 1: Max Calories:", res[len(res)-1])
	fmt.Println("Part 2: Max Calories:", tot)
}
