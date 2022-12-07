package day4

import "fmt"

func part1(a1, b1, a2, b2 int) int {
	ret := 0
	if a1 <= a2 && b1 >= b2 || a1 >= a2 && b1 <= b2 {
		ret++
	}
	return ret
}
func part2(a1, b1, a2, b2 int) int {
	ret := 0
	if a1 <= a2 && b1 >= b2 ||
		a1 >= a2 && b1 <= b2 ||
		a1 >= a2 && a1 <= b2 ||
		b1 >= a2 && b1 <= b2 ||
		a2 >= a1 && a2 <= b1 ||
		b2 >= a1 && b2 <= b1 {
		fmt.Printf("%d-%d,%d-%d overlap\n", a1, b1, a2, b2)
		ret++
	} else {
		fmt.Printf("%d-%d,%d-%d no overlap\n", a1, b1, a2, b2)
	}
	return ret
}

func Solve(lines []string) {
	p1 := 0
	p2 := 0
	for _, line := range lines {
		var a1, b1, a2, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &b1, &a2, &b2)
		p1 += part1(a1, b1, a2, b2)
		p2 += part2(a1, b1, a2, b2)
	}
	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
}
