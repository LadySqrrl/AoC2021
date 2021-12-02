// Package main is the main fucntion
package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	strList, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	count := 1

	for i := range strList {

		if i > 0 {
			if strList[i] > strList[i-1] {
				count++
			}
		}

	}

	fmt.Println(count)
}
