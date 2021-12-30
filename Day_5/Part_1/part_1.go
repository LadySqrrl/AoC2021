// Package main is the main fucntion
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	lineCrossings := [10][10]string{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			lineCrossings[i][j] = "."
		}
	}

	for i := 0; i < len(strList); i++ {
		dir := strings.Split(strList[i], " ")
	}

	fmt.Println(lineCrossings)

}
