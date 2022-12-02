package util

import (
	"bufio"
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInput(day string) *bufio.Scanner {
	path, err := os.Getwd()
	Check(err)

	dfile := fmt.Sprintf("%s/data/day%s.txt", path, day)
	fmt.Println("Loading input for Day", day, "from", dfile)
	fin, err := os.Open(dfile)
	Check(err)

	return bufio.NewScanner(fin)
}
