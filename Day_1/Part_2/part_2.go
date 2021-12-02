// Package main is the main fucntion
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	prevSum := 0
	currSum := 0

	count := -1

	for i := 0; i < len(strList)-2; i++ {

		a, err := strconv.Atoi(strList[i])
		if err != nil {
			fmt.Println("Error converting string to int")
		}
		b, err := strconv.Atoi(strList[i+1])
		if err != nil {
			fmt.Println("Error converting string to int")
		}
		c, err := strconv.Atoi(strList[i+2])
		if err != nil {
			fmt.Println("Error converting string to int")
		}

		currSum = a + b + c

		if currSum > prevSum {
			count++
		}

		prevSum = currSum

	}

	fmt.Println(count)
}
