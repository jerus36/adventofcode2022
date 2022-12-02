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

func getDataPath(day string) string {
	path, err := os.Getwd()
	Check(err)

	return fmt.Sprintf("%s/data/day%s.txt", path, day)
}

func GetInput(day string) *bufio.Scanner {
	dfile := getDataPath(day)
	fmt.Println("Loading input for Day", day, "from", dfile)
	fin, err := os.Open(dfile)
	Check(err)

	return bufio.NewScanner(fin)
}

func ReadLines(day string) []string {
	dfile := getDataPath((day))
	fd, err := os.Open(dfile)
	Check(err)
	defer fd.Close()

	var ret []string
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}
