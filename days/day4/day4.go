package day4

import "fmt"

func part1(a1, b1, a2, b2 int) int {
	fmt.Println(a1, b1, a2, b2)
	return 1
}
func Solve(lines []string) {
	p1 := 0
	for _, line := range lines {
		var a1, b1, a2, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &b1, &a2, &b2)
		p1 += part1(a1, b1, a2, b2)
	}
	fmt.Println(p1)
}
