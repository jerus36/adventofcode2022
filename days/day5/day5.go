package day5

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func buildStacks(lines []string) [9]stack.Stack {

	var ret [9]stack.Stack
	var boxes stack.Stack
	for _, line := range lines {
		if len(line) == 0 {
			break
		} else {
			boxes.Push(line)
		}
	}
	boxes.Pop()
	for boxes.Len() > 0 {
		str := fmt.Sprintf("%v", boxes.Pop())
		for _, r := range str {

		}

		s := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(str, "[", ""), "]", ""), "  ", " -")
		l := strings.Split(s, " ")
		fmt.Println(l)
		for x, v := range l {
			fmt.Println(x, v)
			if v != "" {
				ret[x].Push(v)
			}
		}
	}
	return ret
}
func Solve(lines []string) {
	stacks := buildStacks(lines)
	fmt.Println(stacks)
}
